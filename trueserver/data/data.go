package data

import (
	"fmt"//fmt.Spintfを使うためにインポート
	"math"
) 

type Member struct{
	Name string
	Point int
	Coeff float64
}

//ポイントの値を少数化して、Coeffをかけて返すメソッド.
func Effective(m Member)float64{
	return float64(m.Point)*m.Coeff
}

//作成した値を文字列にして返すメソッド.
func Describe(m Member)string{
	return fmt.Sprintf("%sさんのポイントは%d点、有効ポイントは%.2f点",
	m.Name, m.Point, Effective(m))
}

//上と同じ：別の書き方
//インスタンスをメソッド名の前に置く(この時の引数の事を「レシーバ」と呼ぶ)
func (m Member)EffectiveM() float64{
	return float64(m.Point)*m.Coeff
}

func (m Member)DescribeM() string{
	return fmt.Sprintf("%sさんのポイントは%d点、有効ポイントは%.2f点",
	m.Name, m.Point, m.EffectiveM()) //引数が不要になった
}

//メンバー全体の中で一番ポイントが多い人を返す
func MaxPointMember(members []Member)Member{
	mpm := members[0]

	for _, v := range members{
		if Effective(v) > Effective(mpm) {
			mpm = v
		}
	}

	return mpm
}

//ポインタを利用して、Memberの値を書き換える
func AddPoint(member **Member,p int){
	(**member).Point +=p
}

//上記のメソッドバージョン
func (member *Member)AddPointM(p int){
	member.Point += p
}

//戻り値を返す場合は、本体の書き換えをするわけでないのでポインタ使わなくてOK
func CreateFriendMember(member Member, name string)Member{
	member.Name = name
	return member
}


//*************************************************************************
//旅人タイプ
type Traveller struct{
	Name string
	X int 
	Y int
	Record string
}

//旅人の新規作成
func CreateTraveller(name string, x int, y int)Traveller{
	new_traveller := Traveller{}
	new_traveller.Name = name
	new_traveller.X = x
	new_traveller.Y = y
	new_traveller.Record = fmt.Sprintf("%sさん(%d,%d)地点よりスタート\n",new_traveller.Name, new_traveller.X, new_traveller.Y)
	return new_traveller
}

//旅人の移動(インスタンスで設定しているので、本体が書き換わる)
func(t Traveller)Travel(x int, y int)Traveller{
	t.X = x
	t.Y = y
	t.Record += fmt.Sprintf("(%d,%d)へ移動\n", x, y)
	return t
}

//旅人の到着(インスタンスで設定しているので、本体が書き換わる)
func(t Traveller)Goal()Traveller{
	t.Record += "到着です。\n"
	return t
}

//*************************************************************************
//type Half struct{value float64}と同じ
type Half float64

//type Full struct{value int}と同じ
type Full int

//FractionはValueのメソッドを持っている型
type Fraction interface{
	Value() string
}

//型を整形して返す
func(h Half)Value()string{
	return fmt.Sprintf("%.1f", float64(h))
}

func(f Full)Value()string{
	return fmt.Sprintf("%d",int(f))
}

//*************************************************************************
type Counter interface{
	DoCount()string
}

//文字数を数えたい文字
type CharCounter struct{
	Content string
}

//桁素を数えたい整数
type DigitCounter struct{
	Content int
}

//引数に数えたい文字は取らず、CharCounterに任す
func(counter CharCounter)DoCount()string{

	//処理する値を取り出す
	content := counter.Content
	s_string := fmt.Sprintf("%sは",content)

	//文字列をUnicode記号の配列に変換する([]rune), lenで配列数を取得
	s_string += fmt.Sprintf("%d文字です",len([]rune(content)))

	return s_string
}

//引数に数えたい整数は取らず、DigitCounterに任す
func(counter DigitCounter)DoCount()string{

	//処理する値を取り出す
	content := counter.Content
	content_str := fmt.Sprintf("%d",content) //文字列に変換
	s_string := fmt.Sprintf("%dは",content)

	//文字列をUnicode記号の配列に変換する([]rune), lenで配列数を取得
	s_string += fmt.Sprintf("%d文字です",len([]rune(content_str)))

	return s_string
}

//*************************************************************************
type MockReader interface{
	Read(content string)
	Write() string
}

type StringReader struct{
	Memory string
}

type IntReader struct{
	Memory []int
}

//戻り値を利用せず、ポインタで本体を書き換える
func(reader *StringReader)Read(content string){
	reader.Memory += content
	reader.Memory += "\n"
}

//戻り値を利用せず、ポインタで本体を書き換える
func(reader *IntReader)Read(content string){
	digits := "0123456789"

	//reader.Memory = 数字の配列
	//受け取った配列の文字が数字だった場合のみpushされる
	for _, v := range content{
		for i, s := range digits{
			if v == s{
				//reader.Memoryにインデックス番号をpush
				reader.Memory = append(reader.Memory, i)
			}
		}
	}
}

//受け取った値の型がStringReader型だったらこちらを通る
func(reader StringReader)Write()string{
	s_string := "StringReaderインスタンスの中身は\n"
	s_string += "「"
	s_string += reader.Memory
	s_string += "」"
	return s_string
}

//受け取った値の型がIntReader型だったらこちらを通る
func(reader IntReader)Write()string{
	s_string := "IntReaderインスタンスの中身は\n"
	s_string += "["
	for _, v := range reader.Memory{
		s_string += fmt.Sprintf("%d",v)
	}
	s_string += "]"
	return s_string
}

//*************************************************************************
//インターフェースを使わないメソッド
func(reader IntReader)Reader2Int()int{
	sum := 0
	memory := reader.Memory
	lm := len(memory)
	for i := 0; i<lm; i++{
		mag := math.Pow10(lm-i)
		sum += memory[i]*int(mag)
	}
	return sum/10
}