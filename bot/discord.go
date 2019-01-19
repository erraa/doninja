package bot

import (
	"io"
	"os"
	"strings"

	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/erraa/doninja/config"
	rand "github.com/erraa/doninja/random"
	"github.com/erraa/doninja/utils"
)

var botid string
var goBot *discordgo.Session
var log = utils.LogWithPrefix("Bot")
var conf = config.ReadConfig()

func Start() {
	goBot, err := discordgo.New("Bot " + conf.Discord.Token)
	// Logging

	if err != nil {
		log.Error(err.Error())
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		log.Error(err.Error())
	}

	botid = user.ID

	goBot.AddHandler(imageHandler)
	err = goBot.Open()
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info("Bot is running with ID", botid)

}

func imageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, conf.Discord.BotPrefix) {
		return
	}
	if m.Author.ID == botid {
		return
	}
	if !isValidUrl(m.Content) {
		return
	}
	log.Info("Valid url ", m.Content)
	u, err := url.Parse(m.Content)

	if err != nil {
		log.Warning("Failed to parse ", m.Content)
	}

	pathSlice := strings.Split(u.Path, ".")
	pathEnding := pathSlice[len(pathSlice)-1]
	log.Info(pathEnding)
	if pathEnding == "jpg" || pathEnding == "png" {
		log.Info(pathEnding)
	}

	file, err := download(u, pathEnding)
	upload(file)
}

func download(u *url.URL, format string) (*os.File, error) {
	response, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	file, err := os.Create("/tmp/" + rand.String(20) + "." + format)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func isValidUrl(testUrl string) bool {
	_, err := url.ParseRequestURI(testUrl)
	if err != nil {
		return false
	} else {
		return true
	}
}
