package lg

import (
	"errors"
	"fmt"
	"net/http"
)

// はじめてのGo言語 DI P168~

// --------↓↓↓アンチパターン↓↓↓--------
// type SimpleDataStore struct {
// 	userData map[string]string
// }
// func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
// 	name, ok := sds.userData[userID]
// 	return name, ok
// }
// func NewSimpleDataStore() SimpleDataStore {
// 	return SimpleDataStore{
// 		userData: map[string]string{
// 			"1": "Fred",
// 			"2": "Marry",
// 			"3": "Pat",
// 		},
// 	}
// }
// func LogOutput(message string) {
// 	fmt.Println(message)
// }

// ユーザーを検索して、挨拶をするビジネスロジックを作成する時にはデータの保存場所が必要
// また、ログを残す必要があるので、ログ出力にも依存する。
// しかし、ＤＳやロガーに別の仕組みを使いたいと思うかも知れないので依存を強制したくない。
// ※LogOutputやSimpleDataStoreを直接渡すと、他のロジックに切り替える時に多数の修正が発生してしまう。
// したがって、何に依存しているかを説明したインターフェースを渡す。
// --------↓↓↓修正版↓↓↓--------

// SDS型の宣言
type SimpleDataStore struct {
	userData map[string]string
}

// CDS型の宣言
type ComplexDataStore struct {
	userData map[string]string
}

// SDS型のメソッド
func (s SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := s.userData[userID]
	return name, ok
}

// CDS型のメソッド
func (c ComplexDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := c.userData[userID]
	return name, ok
}

// DSの抽象(SDS,CDS型のメソッドセットに一致するため暗黙的にSDS型のインターフェースとなる。)
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// ファクトリ関数でインスタンスを生成
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Chris",
			"2": "Marry",
			"3": "John",
		},
	}
}

// LoggerAdapter型の定義
type LoggerAdapter func(message string)

// LoggerAdapter型のメソッドの定義
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// メソッドセットから暗黙的にインターフェースを実装
type Logger interface {
	Log(message string)
}

// LoggerAdapter型の関数の定義
func LogOutput(message string) {
	fmt.Println(message)
}

// ビジネスロジックの構造体の定義
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// インスタンスの定義
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// メソッドの定義１
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("SayHello(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("不明なユーザー")
	}
	return name + "さんこんにちは。", nil
}

// メソッドの定義２
func (sl SimpleLogic) SayGoodBy(userID string) (string, error) {
	sl.l.Log("SayGoodBy(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("不明なユーザー")
	}
	return name + "さんさようなら。", nil
}

// インターフェースの定義
type Logic interface {
	SayHello(userID string) (string, error)
	SayGoodBy(userID string) (string, error)
}

// Controller型の定義
type Controller struct {
	l     Logger
	logic Logic
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayHello内: ")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}
