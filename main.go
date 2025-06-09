package main

// 패키지명을 main으로 바꿔야 실행 가능한 프로그램으로 인식됨
// package Sketchbook 이렇게 되어 있을 경우, 라이브러리(모듈)로 인식됨 (실행X)

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func main() {
	// .env 파일을 읽어서 환경변수로 등록 (실패 시 error(객체) 반환, 성공 시 nil-null을 뜻 함)
	err := godotenv.Load(".env") // 대입연산자 := (짧은 변수 선언), 함수 내부에서 새 변수를 선언 시 사용하고, 타입을 자동으로 유추한다
	if err != nil {
		log.Fatal(err) // 에러 메세지 출력 후 즉시 종료 (fmt.Println() 필요X)
	}

	// 환경변수의 값을 읽어옴
	DISCORD_BOT_TOKEN := os.Getenv("DISCORD_BOT_TOKEN") // ; 필요 없다!
	if DISCORD_BOT_TOKEN == "" {                        // if문에 소괄호 안 쓴다!
		log.Fatal("You need to set DISCORD_BOT_TOKEN")
	}

	// 디스코드 세션 생성 (함수가 여러 값을 동시에 반환할 수 있다)
	Sketchbook, err := discordgo.New("Bot " + DISCORD_BOT_TOKEN) // 새 변수 포함 시 := 사용 가능
	if err != nil {
		log.Fatal("Unable to create discord session: ", err)
	}

	err = Sketchbook.Open() // 대입연산자 =, 이미 선언된 변수일 때 사용 (:= 사용 시 오류)
	if err != nil {
		log.Fatal("Unable to connect discord: ", err)
	}
	// defer: 함수 종료 시 반드시 실행될 정리 작업 예약
	// 일반적으로 리소스를 성공적으로 열었을 때, 바로 아래에 defer를 써서 정리 작업을 예약한다.
	defer Sketchbook.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.") // 메세지 출력

	// 채널이름 := make(chan 타입, 버퍼크기)
	// chan: Go 언어에서 고루틴간에 데이터를 안전하게 주고받기 위해 사용하는 통신 도구
	stop := make(chan os.Signal, 1) // os.Signal(운영체제 "시그널") 타입의 채널을 만든다.
	// signal.Notify(채널, Signal값)
	// Signal값이 오면 채널로 전달하도록 등록
	signal.Notify(stop, os.Interrupt)
	// <-채널 : 채널에서 데이터 읽기 (대기상태)
	// 채널 <- 값 : 채널에 데이터 쓰기
	<-stop

	fmt.Println("Shutting down...")
}
