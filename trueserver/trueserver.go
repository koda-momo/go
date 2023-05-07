package main

import (
	"fmt"
	"net/http"
	"trueserver/functions"
	"trueserver/data"
)

func add(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "*** 初めての関数 ***")
	result := functions.Add(5,7)
	fmt.Fprintf(writer, "5+7=%d",result)
}

func sub(writer http.ResponseWriter, req *http.Request){
	a,b := 5,7
	output, result := functions.Sub(a,b) //返ってきた値がそれぞれoutputとresultに入る

	fmt.Fprintln(writer, "*** 複数の値を戻す関数 ***")
	fmt.Fprintf(writer, output, a, b, result)
}

func with_slices(writer http.ResponseWriter, req *http.Request){
	sl_1 := []int{1, 2, 3, 4}
	fmt.Fprintln(writer, "*** スライスそのものを書き換える ***")

	fmt.Fprintln(writer, "\n*** 処理前の値 ***")
	fmt.Fprintln(writer, "\nsl_1は")
	fmt.Fprintln(writer, sl_1)

	functions.AddAll(sl_1, 9)
	fmt.Fprintln(writer, "\n*** 処理後の値 ***")
	fmt.Fprintln(writer, "\nsl_1は")
	fmt.Fprintln(writer, sl_1)

	fmt.Fprintln(writer, "*** スライスのコピーを書き換える ***")
	sl_2 := functions.AddAndCopy(sl_1, 100)

	//コピーを書き換えているので、本体は書き換わっていない
	fmt.Fprintln(writer, "\nsl_1は")
	fmt.Fprintln(writer, sl_1)
	fmt.Fprintln(writer, "\nsl_2は")
	fmt.Fprintln(writer, sl_2)
}


func with_structs(writer http.ResponseWriter, req *http.Request){
	members := []data.Member{
		data.Member{"ゆみこ",56,1.24},
		data.Member{"トシオ",44,0.98},
		data.Member{"かをる",70,1.02},
	}

	fmt.Fprintln(writer, "*** 構造体(型)を用いた関数 ***")
	fmt.Fprintln(writer, functions.DescribeAllMembers(members))

	fmt.Fprintln(writer, "***構造体を戻す関数 ***")
	fmt.Fprintln(writer, functions.DescribeMaxPointMember(members))
	
	fmt.Fprintln(writer, "\n*** 構造体のポインタを用いる関数 ***")
	member_add := &members[0]
	fmt.Fprintf(writer, functions.AddPointAndReport(&member_add, 12))

	fmt.Fprintln(writer, "\nゆみこの値は変更された")
	fmt.Fprintln(writer, functions.Describe(members[0]))

	friend, s_string := functions.CreateFriendAndReport(members[1],"エミコ")
	fmt.Fprintln(writer, s_string)
	fmt.Fprintln(writer, functions.Describe(friend))

	fmt.Fprintln(writer, "\nエミコは追加されていない")
	fmt.Fprintln(writer, functions.DescribeAllMembers(members))

	//エミコをpushして上書き
	members = append(members, friend)

	fmt.Fprintln(writer,"*** メソッドの使用 ***")
	fmt.Fprintln(writer, functions.DescribeM_AllMembers(members))

	fmt.Fprintln(writer,"\n*** お友達紹介得点 ***")
	fmt.Fprintln(writer, functions.AddPointMAndReport(&members[1],20))

	fmt.Fprintln(writer,"\nトシオはポインタで変更したので値が変わった")
	fmt.Fprintln(writer, functions.Describe(members[1]))
}



func with_pointers(writer http.ResponseWriter, req *http.Request){
	mockmemory := []int{325, 14, 160, 440, 16, 175}

	//index=0から順々に数字を入れていく
	fmt.Fprintln(writer,"*** メモリアドレス指定を模倣する ***")
	fmt.Fprintln(writer, "\nアドレス「0」を指定")
	fmt.Fprintln(writer, functions.DescribeMockStruct(mockmemory, 0))

	//index=3から順々に数字を入れていく
	fmt.Fprintln(writer, "\nアドレス「3」を指定")
	fmt.Fprintln(writer, functions.DescribeMockStruct(mockmemory, 3))

	//&b = 値そのものを渡さず、その場所にある値の処理は引数を通して関数の処理に任せる
	fmt.Fprintln(writer, "\n*** ポインタを使う意味 ***")
	a, b := 10,10
	aa := functions.UpdateOrCopy(a, &b)
	fmt.Fprintf(writer, "a=%d,b=%d,aa=%d", a, b, aa)
}


//*************************************************************************
func with_methods(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer,"*** 連続処理 ***")
	marco := data.CreateTraveller("マルコ", 0, 0)

	//.で繋いでいるメソッドTravel ~ Goalまで全て行った結果をmarcoに代入
	marco = marco.Travel(2,3).
	Travel(12,24).
	Travel(45,78).
	Goal()

	fmt.Fprintln(writer, marco.Record)


	fmt.Fprintln(writer,"*** インターフェースの練習 ***")
	//fractions = [1.5 2 2.5 3 3.5]
	fractions := []data.Fraction{
		data.Half(1.5), data.Full(2), data.Half(2.5), data.Full(3), data.Half(3.5),
	}
	fmt.Fprintln(writer,fractions)
	fmt.Fprintln(writer, functions.ShowFractions(fractions))


	fmt.Fprintln(writer,"*** もっとそれらしいインターフェース ***")
	counters := []data.Counter{
		data.CharCounter{"Let's count!"}, //今回はstruct(構造体)で宣言しているので、{}
		data.CharCounter{"一二三四五六七八九"},
		data.DigitCounter{2500},
		data.DigitCounter{1963061},
		data.CharCounter{"以上!"},
	}

	fmt.Fprintln(writer, functions.CountAll(counters))

	fmt.Fprintln(writer,"*** ポインタもインターフェースで実装できる ***")
	var reader data.MockReader //空で変数宣言

	//MockReader型だけど、インターフェースとしてならStringRaderポインタを設定可能
	reader = &data.StringReader{}
	//書き込み
	reader.Read("2023年5月7日")
	reader.Read("Goのインターフェースを学習した")
	reader.Read("難しかった")
	//書き出し
	fmt.Fprintln(writer, reader.Write())

	// 初期化
	reader = &data.IntReader{}
	//書き込み
	reader.Read("21")
	reader.Read("abc") //読み飛ばされる
	reader.Read("75")
	reader.Read("へ3") //一部読み飛ばされる
	//書き出し
	fmt.Fprintln(writer, reader.Write())


	fmt.Fprintln(writer,"\n*** 実装しないメソッドを使う ***")
	//reader = MockReader型なのを、intReader型に変換(型アサーション)
	int_reader := reader.(*data.IntReader)
	fmt.Fprintln(writer, functions.IntReader2Int(*int_reader))
}

func flows(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "*** while文に相当するfor ***")
	fmt.Fprintln(writer, functions.While10(6))
	fmt.Fprintln(writer, functions.While10(13))
	fmt.Fprintln(writer, functions.While10(10))

	fmt.Fprintln(writer, "\n*** forを用いた無限ループ ***")
	fmt.Fprintln(writer, functions.Forever(3))
	fmt.Fprintln(writer, functions.Forever(10000))

	fmt.Fprintln(writer, "\n*** Switch文 ***")
	for i:=0; i<7; i++{
		fmt.Fprintln(writer,functions.Div3(i))
	}

	fmt.Fprintln(writer, "\n*** 多様な条件のSwitch文 ***")
	for i:=0; i<10; i++{
		fmt.Fprintln(writer,functions.DivBy3(i))
	}
}


func generics(writer http.ResponseWriter, req *http.Request){
	sl_int := []int{0,1,2,3,4}
	sl_str := []string{"花","鳥","風","月","猫","蛙","春"}

	fmt.Fprintln(writer, sl_int)
	fmt.Fprintln(writer, sl_str)

	fmt.Fprintln(writer, "*** ジェネリック ***")
	
	//index3の要素を除外する
	sl_3_int, g_str := functions.RemoveByIndex(sl_int,3)
	fmt.Fprintln(writer, g_str)
	fmt.Fprintln(writer, sl_3_int)

	//index0の要素を除外する
	sl_0_str, g_str := functions.RemoveByIndex(sl_str,0)
	fmt.Fprintln(writer, g_str)
	fmt.Fprintln(writer, sl_0_str)
}


func goroutine(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "*** Goroutineによるスレッド管理 ***")
	
	//すぐに実行
	go func(){
		for i := 0; i<3; i++{
			fmt.Fprintln(writer,functions.Record("Hello",i,100))
		}
		fmt.Fprintln(writer,"Hello完了です。")
	}()

	//上が終わったら実行
	for i := 0; i<3; i++{
		fmt.Fprintln(writer,functions.Record("World",i,100))
	}
	fmt.Fprintln(writer,"World完了です。")
}

func a_channel(writer http.ResponseWriter, req *http.Request){
	ch := make(chan string) //送受信するためのチャンネルデータ作成

	go functions.InChannel("最初はグー",0,200,ch)
	go functions.InChannel("ジャンケンポン",0,100,ch)

	//チャンネルの受信
	//チャンネルに1つの値を送ったとき、それを受け取る相手がいないうちはチャンネルは値を受付けず、待たせる
	//なのでx1には最初に送られてきたものが入り、x2には次に送られてきたものが入る 
	x1 :=<-ch
	x2 :=<-ch

	fmt.Fprintln(writer, x1)
	fmt.Fprintln(writer, x2)
}


func files(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "*** ファイルを読む ***")
	readdata := functions.ReadMyFile("data/readtest.txt")
	fmt.Fprintln(writer,readdata)

	fmt.Fprintln(writer, "\n*** ファイルに書く ***")
	writedata := readdata + "\n40,2.4\n" //書き込むデータ
	fmt.Fprintln(writer,functions.WriteMyFile("data/writetest.txt",writedata))
}


func main(){
	http.HandleFunc("/add",add)
	http.HandleFunc("/sub",sub)
	http.HandleFunc("/with_slice",with_slices)
	http.HandleFunc("/with_struct",with_structs)
	http.HandleFunc("/with_pointer",with_pointers)
	http.HandleFunc("/with_method",with_methods)
	http.HandleFunc("/flow",flows)
	http.HandleFunc("/generic",generics)
	http.HandleFunc("/goroutine",goroutine)
	http.HandleFunc("/a_channel",a_channel)
	http.HandleFunc("/file",files)
	http.ListenAndServe(":8090",nil)
}