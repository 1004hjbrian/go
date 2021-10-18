package decoHandler

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler) // 데코 핸들러 함수타입을 정의 , 3개의 리퀘스트를 받는

type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler //http.Handler구현
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //decohanler가 http데코핸들러 만든다
	self.fn(w, r, self.h) //self를 호출하고 w, r호출하고 self.h를 구현
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  h,
	}
}
