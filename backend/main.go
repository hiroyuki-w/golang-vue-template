package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Db DB
var Db *sqlx.DB

// Staff staff
type Staff struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Deptid int64  `db:"deptid" json:"deptid"`
}

// StaffList list
type StaffList []Staff

var selectList = [3]string{"Goo", "Choki", "Paa"}

//atemp := []Janken{Janken{0, "Goo"}, Janken{1, "Choki"}, Janken{2, "Paa"}}

type Choice struct {
	Index int
	Name  string
}

// Choices ジャンケンの種類
type Choices []Choice

// Player プレイヤー
type Player struct {
	Name       string `json:"name"`
	Choice     int    `json:"choice"`
	ChoiceName string `json:"choice_name"`
}

// SetChoice プライヤーの選択を設定
func (player *Player) SetChoice(choices Choices, name string, choice int) {

	for _, v := range choices {
		if v.Index == choice {
			player.Choice = choice
			player.ChoiceName = v.Name
		}
	}
	if player.Choice == 0 {
		rand.Seed(time.Now().UnixNano())
		c := rand.Intn(3)
		player.Choice = choices[c].Index
		player.ChoiceName = choices[c].Name
	}
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
	} else if ChoiceSelf == 1 && ChoiceOpponent == 2 {
		return win
	} else if ChoiceSelf == 1 && ChoiceOpponent == 3 {
		return lose
	} else if ChoiceSelf == 2 && ChoiceOpponent == 1 {
		return lose
	} else if ChoiceSelf == 2 && ChoiceOpponent == 3 {
		return win
	} else if ChoiceSelf == 3 && ChoiceOpponent == 1 {
		return win
	} else if ChoiceSelf == 3 && ChoiceOpponent == 2 {
		return lose
	}

	return draw

}

func main() {
	http.HandleFunc("/api/result", handler)
	http.HandleFunc("/api/result_db", handlerDb)

	http.ListenAndServe(":3000", nil)
}
func handlerDb(w http.ResponseWriter, r *http.Request) {

	var staffList StaffList
	db, err := sqlx.Open("mysql", "root:root@tcp(db:3306)/docker_local")

	if err != nil {
		log.Fatal(err)
	}

	p := "%" + r.FormValue("word") + "%"
	rows, err := db.Queryx("SELECT * FROM staff WHERE name like ? ", p)
	if err != nil {
		log.Fatal(err)
	}
	var staff Staff
	for rows.Next() {
		//rows.Scanの代わりにrows.StructScanを使う
		err := rows.StructScan(&staff)
		if err != nil {
			log.Fatal(err)
		}
		staffList = append(staffList, staff)
	}
	fmt.Println(staffList)

	res, _ := json.Marshal(staffList)
	fmt.Fprint(w, string(res))

}

func handler(w http.ResponseWriter, r *http.Request) {

	choices := Choices{{Index: 1, Name: "Goo"}, {Index: 2, Name: "Choki"}, {Index: 3, Name: "Paa"}}

	//PlayerNo1 := Player{Name: "自分", Choice: ToInt(r.FormValue("choice"))}
	//PlayerNo2 := Player{Name: "相手", Choice: 0}

	PlayerNo1 := Player{}
	PlayerNo2 := Player{}

	PlayerNo1.SetChoice(choices, "自分", ToInt(r.FormValue("choice")))
	PlayerNo2.SetChoice(choices, "相手", 0)
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
