package main

import (
	"WEB9/cipher"

	"fmt"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) { // 데이터를 송신하는 기본 기능
	//Send data
	sentData = data // SentData에 데이터를 저장
}

// type ZipComponent struct { // 압축 Decorator이며 Decorator이기 때문에 다른 Component를 보유
// 	com Component // ZipComponent는 SendComponent를 가짐
// }

// func (self *ZipComponent) Operator(data string) { // ZipComponent를 통해 데이터를 압축
// 	zipData, err := lzw.Write([]byte(data))
// 	if err != nil {
// 		panic(err)
// 	}
// 	self.com.Operator(string(zipData)) // Operator를 통해 압축된 데이터를 호출
// }

type EncryptComponent struct { // 또 다른 Decorator이기 때문에 다른 Component를 가지게 됨
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) { // EncryptComponent를 통해 데이터를 암호화
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

type DecryptComponent struct { // 복호화(Decrypt)에 대한 Component
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) { // DecryptComponent를 통해 데이터를 복호화
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

// type UnzipComponent struct {
// 	com Component
// }

// func (self *UnzipComponent) Operator(data string) {
// 	unzipData, err := lzw.Read([]byte(data))
// 	if err != nil {
// 		panic(err)
// 	}
// 	self.com.Operator(string(unzipData))
// }

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) { // 데이터의 수신
	recvData = data
}

func main() { // 데이터의 암호화후 송신(Send)
	sender := &EncryptComponent{
		key: "abcde", // abcde의 키 값을 부여하여 데이터에 대한 압축

		com: &SendComponent{},
	}

	sender.Operator("ya") // Hello World라는 데이터를 Operator를 통해 송신

	fmt.Println(sentData) // SentData에 Hello World라는 데이터가 저장

	receiver := &DecryptComponent{
		// 암호화된 데이터에 대한 복호화(DecryptComponent)
		key: "abcde",
		com: &ReadComponent{},
	}

	receiver.Operator(sentData) // 송신한 데이터를 수신하는 Operator
	fmt.Println(recvData)       // 받은 데이터를 출력

}

// 전체 순서 정리
// 1) EncryptComponent의 Operator 함수 호출로 데이터가 암호화
// 2) 암호화된 데이터가 EncryptComponent가 가지고 있는 operator인 ZipComponent를 호출
// 3) ZipComponent의 Operator 함수가 호출 및 데이터 압축
// 4) 데이터는 SendComponent를 통해 ZipComponent가 보유중인 Operator를 호출
// 5) SendComponet의 Operator 함수 호출 후 SentData(또는 Send Data)에 데이터가 저장
// 7) 받은 데이터에 대한 복호화(DecryptComponent)???????????????????????????? 및 출력
