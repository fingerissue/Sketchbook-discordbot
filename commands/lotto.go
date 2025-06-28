package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
)

func handleLotto(s *discordgo.Session, i *discordgo.InteractionCreate) {
	Options := i.ApplicationCommandData().Options
	var mode string
	var numbers string
	inputNumbers := [6]int{}

	for _, option := range Options {
		if option.Name == "mode" {
			mode = option.StringValue()
		}
		if option.Name == "numbers" {
			numbers = option.StringValue()
		}
	}

	if mode == "수동" {
		if numbers == "" {
			replyErrorInteraction(s, i, "⚠️ 6개의 숫자를 공백으로 구분하여, 중복없이 입력해 주세요.")
			return
		}

		temp := strings.Fields(numbers)
		if len(temp) != 6 {
			replyErrorInteraction(s, i, "⚠️ 6개의 숫자를 공백으로 구분하여, 중복없이 입력해 주세요.")
			return
		}

		for index, str := range temp {
			n, err := strconv.Atoi(str)
			if err != nil {
				replyErrorInteraction(s, i, "⚠️ "+str+"은 정수가 아닙니다.")
				return
			}
			if n < 1 || n > 45 {
				replyErrorInteraction(s, i, "⚠️ "+str+"은 범위에서 벗어납니다. 1부터 45까지의 숫자를 입력하세요.")
				return
			}
			for j := 0; j < index; j++ {
				if n == inputNumbers[j] {
					replyErrorInteraction(s, i, "⚠️ 중복된 숫자를 입력하지 마세요.")
					return
				}
			}

			inputNumbers[index] = n
		}
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("🎫 lotto 명령어: %s", mode),
		},
	})

	if err != nil {
		log.Println("Failed to respond to ping: ", err)
		replyErrorFollowup(s, i, "⚠️ 응답 중 오류가 발생했습니다.")
		return
	}
}
