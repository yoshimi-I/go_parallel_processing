package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- string) {
	// 何らかのデータを生成する想定
	time.Sleep(2 * time.Second) // データ生成に時間がかかるシミュレーション
	ch <- "データ生成完了"             // データをチャネルに送信
}

func consume(ch <-chan string) {
	data := <-ch // チャネルからデータを受信
	fmt.Println("受信したデータ:", data)
}

func main() {
	ch := make(chan string) // バッファなしチャネルの作成

	go produce(ch) // データを生成するゴルーチンを起動
	go consume(ch) // データを消費するゴルーチンを起動

	time.Sleep(3 * time.Second) // メインゴルーチンを終了させないための待機
}
