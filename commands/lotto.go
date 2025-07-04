package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func handleLotto(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var (
		mode, numbers string
		inputNumbers  [6]int
		outputNumbers [7]int
	)
	var (
		rank, rankColor    int
		rankmsg, rankEmoji string
	)
	Options := i.ApplicationCommandData().Options

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
	} else if mode == "자동" {
		if numbers != "" {
			replyErrorInteraction(s, i, "⚠️ 자동모드에서는 숫자를 입력할 수 없습니다.")
			return
		}

		count := 0
		for count < 6 {
			jungbok := false
			n := rand.Intn(45) + 1

			for j := 0; j < count; j++ {
				if n == inputNumbers[j] {
					jungbok = true
					break
				}
			}
			if jungbok {
				continue
			}

			inputNumbers[count] = n
			count++
		}
	} else {
		replyErrorInteraction(s, i, "⚠️ 수동/자동 중 하나를 선택하세요.")
		return
	}

	userID := s.State.User.ID
	exists := false

	err := DB.QueryRow("select exists(select 1 from user where user_id = ?)", userID).Scan(&exists)
	if err != nil {
		replyErrorInteraction(s, i, "⚠️ SQL문을 실행하는 중 오류가 발생했습니다.")
		log.Println(err)
		return
	}

	if !exists {
		_, err := DB.Exec("insert into user(user_id, money) values(?, 0)", userID)
		if err != nil {
			replyErrorInteraction(s, i, "⚠️ SQL문을 실행하는 중 오류가 발생했습니다.")
			log.Println(err)
			return
		}
	}

	_, err = DB.Exec("update user set money = money - 1000 where user_id = ?", userID)
	if err != nil {
		replyErrorInteraction(s, i, "⚠️ 정상적으로 로또를 구매하지 못했습니다.")
		log.Println(err)
		return
	}

	output := 0
	for output < 7 {
		jungbok := false
		n := rand.Intn(45) + 1

		for i := 0; i < output; i++ {
			if n == outputNumbers[i] {
				jungbok = true
				break
			}
		}
		if jungbok {
			continue
		}

		outputNumbers[output] = n
		output++
	}

	matchCount := 0
	for _, user := range inputNumbers {
		for i := 0; i < 6; i++ {
			if user == outputNumbers[i] {
				matchCount++
				break
			}
		}
	}

	switch matchCount {
	case 6:
		rank = 1
		rankEmoji = "🎉"
		rankColor = 0xFFD700
		_, err = DB.Exec("update user set money = money + 3750000000 where user_id = ?", userID)
		if err != nil {
			replyErrorInteraction(s, i, "⚠️ 당청금을 수령하는데 문제가 발생했습니다.")
			log.Println(err)
			return
		}
	case 5:
		rank = 3
		rankEmoji = "🥉"
		rankColor = 0xCD7F32
		_, err = DB.Exec("update user set money = money + 25000000 where user_id = ?", userID)
		if err != nil {
			replyErrorInteraction(s, i, "⚠️ 당청금을 수령하는데 문제가 발생했습니다.")
			log.Println(err)
			return
		}

		for _, user := range inputNumbers {
			if user == outputNumbers[6] {
				rank = 2
				rankEmoji = "🥈"
				rankColor = 0xC0C0C0
				_, err = DB.Exec("update user set money = money + 250000000 where user_id = ?", userID)
				if err != nil {
					replyErrorInteraction(s, i, "⚠️ 당청금을 수령하는데 문제가 발생했습니다.")
					log.Println(err)
					return
				}
				break
			}
		}
	case 4:
		rank = 4
		rankEmoji = "🏅"
		rankColor = 0x3498DB
		_, err = DB.Exec("update user set money = money + 50000 where user_id = ?", userID)
		if err != nil {
			replyErrorInteraction(s, i, "⚠️ 당청금을 수령하는데 문제가 발생했습니다.")
			log.Println(err)
			return
		}
	case 3:
		rank = 5
		rankEmoji = "🎊"
		rankColor = 0x2ECC40
		_, err = DB.Exec("update user set money = money + 5000 where user_id = ?", userID)
		if err != nil {
			replyErrorInteraction(s, i, "⚠️ 당청금을 수령하는데 문제가 발생했습니다.")
			log.Println(err)
			return
		}
	default:
		rank = -1
		rankEmoji = "💔"
		rankColor = 0xE74C3C
	}

	if rank == 2 {
		rankmsg = fmt.Sprintf("%d등 당첨! (%d개 + 보너스)", rank, matchCount)
	} else if rank == -1 {
		rankmsg = fmt.Sprintf("꽝! 이 정도면 번호가 님 피하는 거 ㅇㅈ?")
	} else {
		rankmsg = fmt.Sprintf("%d등 당첨! (%d개)", rank, matchCount)
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Type:        discordgo.EmbedTypeRich,
					Title:       fmt.Sprintf("%s 로또 결과", rankEmoji),
					Description: rankmsg,
					Color:       rankColor,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "사용자 번호",
							Value: strings.Trim(fmt.Sprint(inputNumbers), "[]"),
						},
						{
							Name:  "당첨 번호",
							Value: fmt.Sprintf("%s + %d", strings.Trim(fmt.Sprint(outputNumbers[:6]), "[]"), outputNumbers[6]),
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Println("Failed to respond to ping: ", err)
		replyErrorFollowup(s, i, "⚠️ 응답 중 오류가 발생했습니다.")
		return
	}
}
