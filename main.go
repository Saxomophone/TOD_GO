package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var pingRunning bool = true

const errorColor = 0xDE0a26

func main() {
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	jimUserID := os.Getenv("JIM_USER_ID")
	jimChannelID := os.Getenv("JIM_CHANNEL_ID")
	pingRunning = false

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	const prefix string = "!TOD"

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}

		if len(args) < 2 {
			embed := &discordgo.MessageEmbed{
				Title: "No command specified",
				Color: errorColor,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		} else if len(args) == 3 && args[1] == "jim" && args[2] == "stop" {
			pingRunning = false
			embed := &discordgo.MessageEmbed{
				Title: "JimBot has been stopped :(",
				Color: 0x94B1FF,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if m.ChannelID != jimChannelID {
				s.ChannelMessageSendEmbed(jimChannelID, embed)
			} //pings in sent channel and jim channel (if not already in jim channel)

			log.Printf("immediately after: \npingRunning: %v", pingRunning)

		} else if len(args) == 2 && args[1] == "jim" {
			pingRunning = true

		} else if args[1] == "greek" {
			greek_alpahbet := []string{
				"α",
				"β",
				"γ",
				"δ",
				"ε",
				"ζ",
				"η",
				"θ",
				"ι",
				"κ",
				"λ",
				"μ",
				"ν",
				"ξ",
				"ο",
				"π",
				"ρ",
				"σ",
				"τ",
				"υ",
				"φ",
				"χ",
				"ψ",
				"ω",
			}

			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			selection := rng.Intn(len(greek_alpahbet))

			embed := &discordgo.MessageEmbed{
				Title:       "Greek Alphabet",
				Description: greek_alpahbet[selection],
				Color:       0x94B1FF,
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		} else {
			embed := &discordgo.MessageEmbed{
				Title: "Invalid command",
				Color: errorColor,
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		}

		ping(jimUserID, jimChannelID, s, pingRunning)
	})

	sess.Identify.Intents = discordgo.IntentsAll

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Bot is running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt) // Listen for interrupt signals
	<-sc
}

func ping(userID string, channelID string, sess *discordgo.Session, pingRunning bool) {
	for pingRunning {
		log.Printf("pingRunning: %v", pingRunning)
		_, err := sess.ChannelMessageSend(channelID, "<@"+userID+">")

		time.Sleep(500 * time.Millisecond)

		if err != nil {
			fmt.Printf("Failed to send message: %v\n", err)
			break
		}
	}
}
