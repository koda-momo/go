package functions

import (
	"fmt"
	"time"
	"os" //ファイルの読み込み用
	"trueserver/data"
)

//メソッド名は大文字始まり
func Add(a int, b int) int {
	return a + b
}

//戻り値を2つ返す
func Sub(a,b int) (string, int){
	return "%d-%dは%dなのだ", a-b
}

func AddAll(sl []int, a int){
	//len = 要素数(length)
	for i:=0; i<len(sl); i++{
		sl[i] += a
	}
}

func AddAndCopy(sl []int, a int) []int {
	//sl本体を書き換えたくないので、コピーを作成
	sl_cp := []int{}

	for i:=0; i<len(sl); i++{
		//sl_cpにpush(push対象の配列, 加える値)
		sl_cp = append(sl_cp, sl[i]+a)
	}

	return sl_cp
}

//data.goで作成した型を利用したメソッド
func Describe(member data.Member)string{
	s_string := fmt.Sprintf(data.Describe(member))

	s_string += "\n"
	return s_string
}


func DescribeAllMembers(members []data.Member)string{
	s_string := ""

	for _, v := range members {
		s_string += Describe(v)
	}

	return s_string
}

func DescribeM_AllMembers(members []data.Member)string{
	s_string := "メソッドを使って書き出しても\n"
	for _, v := range members {
		s_string += v.DescribeM() //呼び出し方はこの形になる
		s_string += "\n"
	}

	return s_string
}

func DescribeMaxPointMember(members []data.Member)string{
	s_string := "有効ポイントが最大の人は\n"
	mpm := data.MaxPointMember(members)

	s_string += fmt.Sprintf("%sさん", mpm.Name)
	s_string += "\n"

	return s_string
}

//ポインタ
func DescribeMockStruct(mockmemory []int,mockaddress int)string{
	s_string := fmt.Sprintf("名前は%dさん、",mockmemory[mockaddress])
	s_string += fmt.Sprintf("年齢%d歳、",mockmemory[mockaddress + 1])
	s_string += fmt.Sprintf("身長は%dcm",mockmemory[mockaddress + 2])
	return s_string
}

//b *int = 変数bからはint型のデータを読み出します
func UpdateOrCopy(a int, b *int)int{ 
	a += 3

	//bが保持しているアドレスのデータは整数なので、これに3を足せ
	*b += 3

	return a
}

//受け取ったmember情報本体を書き換えるメソッド.
func AddPointAndReport(member **data.Member, p int)string{
	data.AddPoint(member, p)
	s_string := "<<得点アップサービス>>"
	s_string = fmt.Sprintf("%sさんのポイント%d点アップ\n",(**member).Name, p)
	return s_string
}

//上記のメソッドバージョン
func AddPointMAndReport(member *data.Member, p int)string{
	//対象のインスタンスとしてmemberを渡す(インスタンスを設定しているので、*は1つで良い...?)
	member.AddPointM(p)

	//member.Name = (*member).Nameらしい。。(?)
	s_string := "<<メソッドによる得点アップサービス>>"
	s_string = fmt.Sprintf("%sさんのポイント%d点アップ\n",member.Name, p)
	return s_string
}

func CreateFriendAndReport(member data.Member, friend_name string)(data.Member, string){
	friend := data.CreateFriendMember(member, friend_name)
	s_string := fmt.Sprintf("%sさんの紹介で、お友達%sさんが加わりました。",member.Name, friend_name)
	s_string += "\n"
	return friend ,s_string
}


//*************************************************************************
func ShowFractions(fractions[]data.Fraction)string{
	s_string := "スケール的には\n"

	for _, v := range fractions{
		s_string += fmt.Sprintf("%s倍か",v.Value()) //配列の数字(v)をValueにインターフェースとして渡している
		s_string += "\n"
	}

	s_string += "というところでしょうか。\n"
	return s_string
}

//*************************************************************************
//[]Counter型で受け取っている = CharCounter or DigitCounterの値が入った配列
func CountAll(counters []data.Counter)string{
	s_string := "<<data.Counterインターフェース>>\n"

	for _, v := range counters {
		s_string += v.DoCount() //DoCountにvを渡す
		s_string += "\n"
	}

	return s_string
}

//*************************************************************************
func IntReader2Int(reader data.IntReader)string{
	s_string := "IntReaderから構成される整数は\n"
	s_string += fmt.Sprintf("%d", reader.Reader2Int())
	s_string += "\n"
	return s_string
}

//*************************************************************************
//goのwhile文
func While10(num int)string{
	s_string := fmt.Sprintf("最初は%d\n",num)


	if num <10{  //～9
		for num <10{
			num++
			s_string += fmt.Sprintf("%d\n",num)
		}
	}else if num >10{ //{とelse ifは改行不可　11～
		for num >10{
			num--
			s_string += fmt.Sprintf("%d\n",num)
		}
	}else{ //{とelseは改行不可 10の時に通る
		s_string += fmt.Sprintf("今も%d\n",num)
	}
	return s_string
}

//無限ループ
func Forever(limit int)string{
	i := 0
 	for{ //forの後何も書かなければ無限ループ
		i++

		//無限に回るので、止めるタイミングを作っておく
		if i > limit{
			return fmt.Sprintf("%dでやめました。",i)
		}
	}
}

//*************************************************************************
func Div3(num int)string{
	s_string := "3は"
	switch num{
	case 0:
		s_string += "0では割れません。"
	case 1:
		s_string += "1で割る意味はあまりない"
	case 2:
		s_string += "2で割ると1と1/2"
	case 3:
		s_string += "3で割るとちょうど1"
	default:
		if num%3==0{
			s_string += fmt.Sprintf("%dで割ると1/%d",num,num/3)
		}else{
			s_string += fmt.Sprintf("%dで割ると3/%d",num,num)
		}
	}
	return s_string
}

func DivBy3(num int)string{
	s_string := fmt.Sprintf("%dを3で割る",num)
	m := num%3
	switch {
	case num <1:
		s_string += "のは考えない"
	case num <3:
		s_string += fmt.Sprintf("と%d/3",num)
	case num ==0:
		s_string += fmt.Sprintf("と%d",num/3)
	default:
		s_string += fmt.Sprintf("と%dと%d/3",(num-m)/3,m)
	}
	return s_string
}


//*************************************************************************
//ジェネリック型:sl=Tの配列, index=intを受け取って、Tの配列とstringを返す
func RemoveByIndex[T any](sl []T, index int)([]T, string){

	s_string := fmt.Sprintf("インデックスが%dの要素を除きます",index)
	rest := []T{} //引数Tの型に合わせた空配列

	for i,v := range sl{
		if i != index { //除きたいindex番号でなければ、restに値をpush
			rest = append(rest, v)
		}
	}

	return rest, s_string
}

//*************************************************************************
//並行処理
func Record(s string, times int, interval int)string{
	time.Sleep(time.Duration(interval)*time.Millisecond)
	return fmt.Sprintf("%s_%d",s,times)
}

//チャンネル
func InChannel(s string, times int, interval int, ch chan string){ //ch = chan:string型
	//チャンネルの送信
	ch<-Record(s, times, interval)
}

//*************************************************************************
//ファイルの読み込み
func ReadMyFile(filepath string)string{
	data, err := os.ReadFile(filepath)
	if err != nil{ //エラーに値が入っていたらエラーの表示
		return "ファイルを開けませんでした。"
	}
	return string(data)
}

//ファイルの書き込み
func WriteMyFile(filepath string, content string)string{
	err := os.WriteFile(filepath, []byte(content), 0666) //0666=ファイルのアクセス権限の定数
	if err != nil{ //エラーに値が入っていたらエラーの表示
		return "ファイルを書き込めませんでした。"
	}
	return fmt.Sprintf("%sに書き込みました。",filepath)
}