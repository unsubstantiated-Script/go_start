package my_logger

import "log"

func ListenForLog(ch chan string) {
	for {
		msg := <-ch
		log.Println(msg)
	}
}
