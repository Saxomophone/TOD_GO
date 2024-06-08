package main

import (
    "github.com/bwmarrin/discordgo"
    "log"
    "fmt"
    "os"
    "os/signal"
    "syscall"
		"strings"
		"math/rand"
		"time"
		"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	sess, err := discordgo.New("Bot " + token)
	if err!= nil {
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
		
		if args[1] == "greek" {
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
		
		rand.Seed(time.Now().UnixNano())
		selection := rand.Intn(len(greek_alpahbet))


		embed := &discordgo.MessageEmbed{
			Title: "Greek Alphabet",
			Description: greek_alpahbet[selection],
			Color: 0x94B1FF,
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAll

	err = sess.Open()
	if err!= nil {
		log.Fatal(err)
	}
	defer sess.Close()


	fmt.Println("Bot is running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt) // Listen for interrupt signals
	<-sc
}
