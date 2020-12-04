package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

/*type Config struct {
	storage store.Store
}

var config Config
*/
// https://media1.tenor.com/images/baf2d324d696b8e0b08daa8cff5c8f12/tenor.gif?itemid=12992329
func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Errorf("error initializing discord bot")
	}
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	dg.AddHandler(createOrder)

	err = dg.Open()

	if err != nil {
		fmt.Errorf("error opening connection")
		return
	}

	fmt.Println("Bot is running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func createOrder(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!ayaya" {
		if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
			fmt.Printf("Error deleting message\n")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "https://media1.tenor.com/images/baf2d324d696b8e0b08daa8cff5c8f12/tenor.gif?itemid=12992329")
	} else if m.Content == "!uiuiui" {
		if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
			fmt.Printf("Error deleting message\n")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "https://media.tenor.com/images/a074d3d12d981c5876d4f08daddf9b5a/tenor.gif")
	}

}
