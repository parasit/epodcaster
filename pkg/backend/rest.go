package backend

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/ermanimer/log/v2"

	logging "github.com/parasit/epodcaster/pkg/log"
	models "github.com/parasit/epodcaster/pkg/models"
	storage "github.com/parasit/epodcaster/pkg/storage"
	//tools "github.com/parasit/epodcaster/pkg/tools"
)

var l *log.Logger
var c models.Channel

func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	logging.Log.Debug("Get index")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(storage.GetChannels())
}

func addChannel(w http.ResponseWriter, r *http.Request) {
	addCorsHeader(w)

	logging.Log.Debug("Add channel (", r.Method, ")")

	switch r.Method {
	case "OPTIONS":
		logging.Log.Debug("AddChannel OPTIONS")
		w.WriteHeader(http.StatusOK)
		return
	case "GET":
		logging.Log.Debug("AddChannel GET")
		http.ServeFile(w, r, "index.html")
		return
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		body, _ := ioutil.ReadAll(r.Body)
		logging.Log.Debug("Post body ", string(body))
		var tChannel models.Channel
		err := json.Unmarshal(body, &tChannel)
		if err != nil {
			w.Write([]byte("{'response':'wrong data'}"))
			w.WriteHeader(http.StatusOK)
		}
		//w.WriteHeader(http.StatusOK)
		storage.AddChannel(tChannel)
		w.Write([]byte("{\"response\":\"ok\"}"))
		logging.Log.Debug(tChannel.Author)
	}
}

func deleteChannel(w http.ResponseWriter, r *http.Request) {
	channelID, err := strconv.Atoi(r.URL.Query().Get("channel"))
	logging.Log.Debug("Delete channel ", channelID)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	storage.DeleteChannel(channelID)
	w.WriteHeader(http.StatusOK)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
  addCorsHeader(w)
	var n, fsize int
	m, p, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	logging.Log.Debug(m)
	boundary := p["boundary"]
	reader := multipart.NewReader(r.Body, boundary)
	buff := make([]byte, 1024)
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			// Done reading body
			break
		}
		contentType := part.Header.Get("Content-Type")
		logging.Log.Debug(contentType)
		fname := part.FileName()
		logging.Log.Debug("Filename %s", fname)
		logging.Log.Debug("Open file")
		f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		for {
			_, err = part.Read(buff)
			if err == io.EOF {
				break
			}
			if err != nil {
				logging.Log.Fatal(err)
			}
			n, err = f.Write(buff)
			fsize += n
			if err != nil {
				panic(err)
			}
		}
		logging.Log.Debug("wrote ", fsize, " bytes")
	}
}

// Run main backend process
func Run() {
	// c.LoadFromFile("test.json")
	// logging.Log.Debug("Data loaded")
	// c := storage.FindChannelById(1)
	// episodes := tools.CheckFolder("vm1")
	// tools.ParseEpisodes(c, episodes)
	// logging.Log.Debug("Episodes parsed")
	http.HandleFunc("/", getIndex)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/addchannel", addChannel)
	http.HandleFunc("/deletechannel", deleteChannel)

	logging.Log.Error(http.ListenAndServe(":8080", nil))
}
