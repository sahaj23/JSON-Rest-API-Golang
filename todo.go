package main

import (
	// "fmt"
	thread "go-rest-api/thread"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mux sync.RWMutex

var Threads thread.ThreadSlice = thread.ThreadSlice{{Id: 1, Discription: "My Discription", Title: "My title", TimeCreated: "15:19:19, Jun 10 2019"},
	{Id: 2, Discription: "My Discription2", Title: "My title2", TimeCreated: "15:19:19, Jun 10 2019"}}

func getAllThreads(writestr http.ResponseWriter) {
	jsonbyte, err := Threads.MarshalJSON()
	if err != nil {
		writestr.Write([]byte("There was a Problem Processing your request"))
		return
	}

	writestr.Write(jsonbyte)
}
func getThreadFromId(id int, writestr http.ResponseWriter) {
	for _, val := range Threads {
		if val.Id == id {
			jsonbyte, err := val.MarshalJSON()
			if err != nil {
				writestr.Write([]byte("There was a Problem Processing your request"))
				return
			}
			writestr.Write(jsonbyte)
			return

		}
	}
	writestr.Write([]byte("Id not Found"))

}
func addThread(newThread thread.Thread, writestr http.ResponseWriter) {

	newThreadId := newThread.Id
	for _, val := range Threads {

		if val.Id == newThreadId {
			writestr.Write([]byte("Id Already Exists"))
			return
		}
	}
	currentTimeString := time.Now().Format(" 15:04:05, Jan _2 2006")

	myThread := thread.Thread{Id: newThreadId,
		Discription: newThread.Discription,
		Title:       newThread.Title,
		TimeCreated: currentTimeString}

	Threads = append(Threads, myThread)
	writestr.Write([]byte("Thread Added Successfully"))
	return

}
func DeleteThread(id int, writestr http.ResponseWriter) {

	for index, val := range Threads {
		if val.Id == id {
			Threads = append(Threads[:index], Threads[index+1:]...)
			writestr.Write([]byte("Thread Deleted Successfully"))
			return

		}
	}

	writestr.Write([]byte("Id not Found"))
	return
}

func PutThread(newThread thread.Thread, writestr http.ResponseWriter) {

	newThreadId := newThread.Id

	for index, val := range Threads {
		if val.Id == newThreadId {

			timee := Threads[index].TimeCreated
			myThread := thread.Thread{Id: newThreadId,
				Discription: newThread.Discription,
				Title:       newThread.Title,
				TimeCreated: timee}

			Threads[index] = myThread
		}
	}
	writestr.Write([]byte("Thread Changed Successfully"))

	return

}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	pathVars := strings.Split(r.URL.Path, "/")

	if r.Method == "GET" {
		if len(pathVars) > 1 && pathVars[len(pathVars)-1] != "" {
			idStr := pathVars[len(pathVars)-1]

			reqId, err := strconv.Atoi(idStr)
			if err != nil {
				w.Write([]byte("There was an error Processing your request"))
				return
			}
			mux.RLock()
			getThreadFromId(reqId, w)
			mux.RUnlock()
			return
		}
		mux.RLock()

		getAllThreads(w)
		mux.RUnlock()
		return
	}
	if r.Method == http.MethodPost {
		mux.Lock()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("There was an error Processing your request"))
			return
		}

		defer r.Body.Close()

		newThreadData := thread.Thread{}

		erro := newThreadData.UnmarshalJSON(b)

		if erro != nil {
			w.Write([]byte("There was an error Processing your request"))
			return

		}
		// newThreadData:=input["Thread"].(map[string]interface{})
		addThread(newThreadData, w)
		mux.Unlock()
		return
	}
	if r.Method == http.MethodDelete {

		if len(pathVars) > 1 && pathVars[len(pathVars)-1] != "" {
			idStr := pathVars[len(pathVars)-1]

			reqId, err := strconv.Atoi(idStr)
			if err != nil {
				w.Write([]byte("There was an error Processing your request"))
				return
			}
			mux.Lock()
			DeleteThread(reqId, w)
			mux.Unlock()
			return
		}
		w.Write([]byte("There was an error Processing your request"))
		return
	}

	if r.Method == http.MethodPut {
		mux.Lock()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {

			w.Write([]byte("There was an error Processing your request"))
			return
		}

		defer r.Body.Close()
		newThreadData := thread.Thread{}
		erro := newThreadData.UnmarshalJSON(b)
		if erro != nil {
			w.Write([]byte("There was an error Processing your request"))
			return
		}
		PutThread(newThreadData, w)
		mux.Unlock()
	}

	return

}

func main() {

	http.HandleFunc("/Threads/", HandleRequest)
	http.ListenAndServe(":8080", nil)
}
