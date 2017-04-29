package matcher

import (
	"fmt"
	"github.com/cset"
	"github.com/pborman/uuid"
	"log"
	"net/http"
)

type Pair struct {
	Male            string
	Female          string
	MaleConnected   bool
	FemaleConnected bool
}

var males = cset.NewSet()
var females = cset.NewSet()

var paidMales = cset.NewSet()
var paidFemales = cset.NewSet()
var paired = cset.NewMap()
var roomMap = cset.NewMap()

func AddMale(name string) {
	males.Add(name)
}

func RemoveMale(name string) {
	males.Remove(name)
}

func AddFemale(name string) {
	females.Add(name)
}

func RemoveFemale(name string) {
	females.Remove(name)
}

func HandlePaidMale(name string) {
	if males.Has(name) && !paidMales.Has(name) {
		paidMales.Add(name)
		males.Remove(name)
	}
}

func HandlePaidFemale(name string) {
	if females.Has(name) && !paidFemales.Has(name) {
		paidFemales.Add(name)
		females.Remove(name)
	}
}

// MatchHandler handles a match
func MatchHandler(w http.ResponseWriter, r *http.Request) {
	gender := r.URL.Query().Get("gen")
	uname := r.URL.Query().Get("uname")
	roomID := uuid.NewRandom().String()
	w.Header().Set("Content-Type", "application/json")

	if gender == "m" {
		female := paidFemales.Pop()
		if female == "" {
			female = females.Pop()
		}
		if female == "" {
			fmt.Fprint(w, `{"code":200,"inst":"wait"}`)
		} else {
			paired.Set(female, roomID)
			paired.Set(uname, roomID)
			roomMap.Set(roomID, &Pair{Male: uname, Female: female})
			w.Write([]byte(`{"code":200,"inst":"join","data":"` + roomID + `"}`))
		}
	} else if gender == "f" {
		male := paidMales.Pop()
		if male == "" {
			male = males.Pop()
		}
		if male == "" {
			fmt.Fprint(w, `{"code":200,"inst":"wait"}`)
		} else {
			paired.Set(male, roomID)
			paired.Set(uname, roomID)
			roomMap.Set(roomID, &Pair{Male: male, Female: uname})
			w.Write([]byte(`{"code":200,"inst":"join","data":"` + roomID + `"}`))
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprint(w, `{"code":404,"msg":"Bad Request"}`)
	}
}

// Disconnect disconnects a user from his/her room
func Disconnect(uname string) {
	if roomID, ok := paired.Get(uname).(string); ok {
		pair := roomMap.Get(roomID).(Pair)
		paired.Remove(pair.Male)
		paired.Remove(pair.Female)
		roomMap.Remove(roomID)
	} else {
		log.Println("WARNING:Disconnect:User not found")
	}
}

// MyRoomHandler provides look up method for getting user's current room information
func MyRoomHandler(w http.ResponseWriter, r *http.Request) {

	uname := r.URL.Query().Get("uname")
	w.Header().Set("Content-Type", "application/json")

	if roomID, ok := paired.Get(uname).(string); ok {
		w.Write([]byte(`{"code":200,"data":"` + roomID + `"}`))
	} else {
		w.Write([]byte(`{"code":200,"msg":"Room not found", "data": null}`))
	}
}
