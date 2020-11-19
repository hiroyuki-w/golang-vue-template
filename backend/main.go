package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var selectList = [3]string{"Goo", "Choki", "Paa"}

//atemp := []Janken{Janken{0, "Goo"}, Janken{1, "Choki"}, Janken{2, "Paa"}}

// Choices ジャンケンの種類
type Choices []struct {
	Index int
	Name  string
}

// Player プレイヤー
type Player struct {
	Name   string `json:"name"`
	Choice int    `json:"choice"`
}

// SetChoice プライヤーの選択を設定
func (player *Player) SetChoice(choices Choices) {
	rand.Seed(time.Now().UnixNano())
	player.Choice = choices[rand.Intn(3)].Index
}

// Judgement 勝敗判定
type Judgement struct {
	PlayerSelf     Player
	PlayerOpponent Player
	Choices        Choices
}

// Response 勝敗レスポンス
type Response struct {
	PlayerSelf     Player
	PlayerOpponent Player
	Result         int `json:"result"`
}

// GetJudge 勝敗取得
func (Judge Judgement) GetJudge() int {
	ChoiceSelf := Judge.PlayerSelf.Choice
	ChoiceOpponent := Judge.PlayerOpponent.Choice

	win := 1
	lose := 2
	draw := 0

	if ChoiceSelf == ChoiceOpponent {
		return draw
	} else if ChoiceSelf == 0 && ChoiceOpponent == 1 {
		return win
	} else if ChoiceSelf == 0 && ChoiceOpponent == 2 {
		return lose
	} else if ChoiceSelf == 1 && ChoiceOpponent == 0 {
		return lose
	} else if ChoiceSelf == 1 && ChoiceOpponent == 2 {
		return win
	} else if ChoiceSelf == 2 && ChoiceOpponent == 0 {
		return win
	} else if ChoiceSelf == 2 && ChoiceOpponent == 1 {
		return lose
	}

	return draw

}

func main() {
	http.HandleFunc("/api/fuga", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	choices := Choices{{Index: 0, Name: "Goo"}, {Index: 1, Name: "Choki"}, {Index: 2, Name: "Paa"}}

	PlayerNo1 := Player{Name: "自分", Choice: 0}
	PlayerNo2 := Player{Name: "相手", Choice: 0}

	PlayerNo1.SetChoice(choices)
	PlayerNo2.SetChoice(choices)
	//fmt.Fprint(w, PlayerNo1)
	//fmt.Fprint(w, PlayerNo2)
	Judgement := Judgement{PlayerNo1, PlayerNo2, choices}

	result := Judgement.GetJudge()

	//:= Response{PlayerSelf: PlayerNo1, PlayerOpponent: PlayerNo2, PlayerWin: WinPlayer}

	//json, _ := json.Marshal(Response{PlayerSelf: PlayerNo1, PlayerOpponent: PlayerNo2, PlayerWin: WinPlayer})
	tmp := Response{PlayerSelf: PlayerNo1, PlayerOpponent: PlayerNo2, Result: result}
	tmp2, _ := json.Marshal(tmp)
	fmt.Fprint(w, string(tmp2))
	//a, _ := json.Marshal(WinPlayer)

	//fmt.Fprint(w, string(a))
	//fmt.Fprint(w, string(json))

}

// ToInt int型に変換
func ToInt(str string) int {
	integer, _ := strconv.Atoi(str)
	return integer
}
