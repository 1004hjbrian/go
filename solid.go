// SOLID 원칙
//단일 책임의 srp
//함수는확장수정은닫혀있는 ocp
//객체는 하위 타입의 인스턴스로 바꿀수 있는 ISP
//특정클라이언트를 위한 하나 인터페이스보다 여러개가 범용 인터페이스보다 낫다LSP
//구체화x추상화에 의존 DIP

package main

import (
	"fmt"
)

// 인터페이스 1 - 부산 회. 얘가 앞으로 SpoonOfJam 역할을 할꺼에요 / 한개의 type 문 아래  기능을 추가하려면 기존 func의 코드를 건드려야 한다.
// 그것은 단일 책임 srp ,  수정이 닫혀있는 ocp, LSP의 원칙에 위배되므로  여러 인터페이스와 구조체를 만들어 원칙을 준수한다.
type Dish interface {
	String() string
}

// 인터페이스 2 - 회를 젓가락으로 집어요.
type Slice interface {
	getonedish() Dish
}

type Sashimi struct {
	val string
}

//  Fillet 회 뜨다.
func (s *Sashimi) Fillet(Slice Slice) {
	pickit := Slice.getonedish()
	s.val += pickit.String()
}

func (s *Sashimi) String() string {
	return "(( 멍게 " + s.val
}

// 모듬회에 따라나오는 서비스를 통합해서 새로운 interface로 정의해요.
// 인터페이스 3 - 서비스
type Service interface {
	String() string
}

// 인터페이스 4 - 사장님의 마음이 담긴 무료 서비스
type Free interface {
	HeartofOwner() Service
}

type MarketService struct {
	val string
}

func (f *MarketService) pickService(Free Free) {
	serve := Free.HeartofOwner()
	f.val += serve.String()
}

func (f *MarketService) String() string {
	return f.val
}

// ----- 손질된 회를 접시에 올립니다 ------- //

// 광어를 회 떠요
type FlatfishSlice struct {
}

func (j *FlatfishSlice) getonedish() Dish {
	return &PieceofFlatfishSlice{}
}

type PieceofFlatfishSlice struct {
}

func (s *PieceofFlatfishSlice) String() string {
	return " + 광어 한 마리 "
}

// 우럭을 회 떠요
type RockfishSlice struct {
}

func (j *RockfishSlice) getonedish() Dish {
	return &PieceofRockfishSlice{}
}

type PieceofRockfishSlice struct {
}

func (s *PieceofRockfishSlice) String() string {
	return " + 우럭 한 마리 "
}

// 도미를 회 떠요
type BreamSlice struct {
}

func (j *BreamSlice) getonedish() Dish {
	return &PieceofBreamSlice{}
}

type PieceofBreamSlice struct {
}

func (s *PieceofBreamSlice) String() string {
	return " + 도미 한 마리"
}

// 우럭을 회 떠요
type BassSlice struct {
}

func (j *BassSlice) getonedish() Dish {
	return &PieceofBassSlice{}
}

type PieceofBassSlice struct {
}

func (s *PieceofBassSlice) String() string {
	return " + 농어 한 마리"
}

// ---------- 서비스 인터페이스 ---------- //

// 서비스 1번 대구탕

type SpicyFishStew struct {
}

func (f *SpicyFishStew) HeartofOwner() Service {
	return &ServiceSpicyFishStew{}
}

type ServiceSpicyFishStew struct {
}

func (s *ServiceSpicyFishStew) String() string {
	return "생선머리 넣고 끓인 대구탕"
}

// 서비스 2번 콜라

type Cola struct {
}

func (f *Cola) HeartofOwner() Service {
	return &MarketPepsiCola{}
}

type MarketPepsiCola struct {
}

func (s *MarketPepsiCola) String() string {
	return " + 업소용 콜라 1.5L"
}

func main() {
	// 회와 서비스 주소를 정의해볼께요
	Sashimi := &Sashimi{}
	MarketService1 := &MarketService{}

	// 개별 횟감과 개별 서비스들의 주소를 정의해볼께요
	Slice := &BassSlice{}
	Slice1 := &FlatfishSlice{}
	Slice2 := &BreamSlice{}
	Slice3 := &RockfishSlice{}

	Service00 := &SpicyFishStew{}
	Service01 := &Cola{}

	// 횟감과 서비스를 손님계신 자리로 서빙할께요
	Sashimi.Fillet(Slice)
	Sashimi.Fillet(Slice1)
	Sashimi.Fillet(Slice2)
	Sashimi.Fillet(Slice3)

	MarketService1.pickService(Service00)
	MarketService1.pickService(Service01)

	// 완성!
	fmt.Println(" *:･｡,☆ﾟ’･:*:･*:･｡,☆ﾟ’･:*:･｡환    영｡･:*:･ﾟ’☆,｡･:**:･｡,☆ﾟ’･:*:･ ")
	fmt.Println("----------------만두네 아낌없이주는횟집 부산해운대점------------------")
	fmt.Println("")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~~ 천사채 ~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println(Sashimi, "))")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~~ 천사채 ~~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println("")
	fmt.Println("서비스 메뉴 : ", MarketService1, " 나왔습니다!")
	fmt.Println("")
	fmt.Println("----------------주문하신 부산 해운대 모듬회 나왔습니다-----------------")
	fmt.Println("----------------부족하시면 알아서 잡아다 회 떠 드세요-----------------")
	fmt.Println("*:･｡,☆ﾟ’･:*:･*:･｡,☆ﾟ’･:*:･｡안녕히 가십시오｡･:*:･ﾟ’☆,｡･:**:･｡,☆ﾟ’･:*:･")
}
