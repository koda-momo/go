package main

import(
	"fmt"
	"net/http" //webサーバーとして使用する為のパッケージ
	"math"
)

func hello(writer http.ResponseWriter, req *http.Request){
 fmt.Fprintln(writer, "レッツゴー\n")
}

func algebra(writer http.ResponseWriter, req *http.Request){
	a, b := 7,8 //代入しつつ、型を自動判別してほしい場合は:=
	result1 := "%d+%d=%d\n"
	fmt.Fprintln(writer, "*** 整数の足し算 ***")
	fmt.Fprintf(writer, result1, a, b, a+b) //(writer, 変数名, 変数のd一つ目に入れる値...)

	c := 15.0 
	d := float64(a)
	result2 := "%.1f/%.1f=%.3f\n" //.1f小数点第一位, .3f小数点第三位
	fmt.Fprintln(writer, "\n*** 小数の割り算 ***")
	fmt.Fprintf(writer, result2, c, d, c/d)
}

func arrays(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "*** 要素が五個の配列を定義 ***")
	arr1:=[5]int{2,4,6,8,10}
	fmt.Fprintln(writer, arr1)

	fmt.Fprintln(writer, "\n*** 要素が五個の配列で3つしか入れない場合 ***")
	arr2:=[5]int{2,4,6}
	fmt.Fprintln(writer, arr2)

	fmt.Fprintln(writer, "\n*** 要素の値を変更する ***")
	arr2[4]=99
	fmt.Fprintln(writer, arr2)

	fmt.Fprintln(writer, "\n*** slice ***")
	slice1 := arr1[1:3] //1-2箱目を取得(1-3未満)
	slice2 := arr2[3:] //3箱目以降を取得
	fmt.Fprintln(writer, slice1)
	fmt.Fprintln(writer, slice2)

	fmt.Fprintln(writer, "\n*** sliceした値を変更したら元の値も変わるか? ***")
	slice1[1] = 36
	fmt.Fprintln(writer, slice1)
	fmt.Fprintln(writer, arr1)
}

func slices(writer http.ResponseWriter, req *http.Request){
	sl := []int{30, 45, 60, 90, 180}
	fmt.Fprintln(writer, sl)

	var rad_v float64 //初期値を定めたくない場合の変数宣言のやり方

	//_=インデックス, v=slの中身１つずつ, for + rangeで回せる
	for _, v := range sl{
		//空の変数rad_vに、順番に少数化したslの中身 * math.Pi/180.0を入れて行く
		//math.Pi = mathパッケージのPi(円周率)
		rad_v = float64(v)*math.Pi/180.0
		fmt.Fprintf(writer, "\nsin%d°は %.3f\n\n", v, math.Sin(rad_v))
	}
   }

   
//member型
type member struct{
	name string
	point int 
	coeff float64
}

//vip型：member型の拡張
type vip struct{
 member
 vip_point int
}

func struct_members(writer http.ResponseWriter, req *http.Request){
 fmt.Fprintln(writer, "*** 構造体memberのインスタンス ***")
 yumiko := member{"ゆみこ", 56, 1.24}
 
 toshio := member{}
 toshio.name = "トシオ"
 toshio.point = 44
 toshio.coeff = 0.98

 members := []member{yumiko, toshio}
 effective := "%sさんの有効ポイントは%.2f\n"

 for _, v := range members{
	fmt.Fprintf(writer, effective, v.name, float64(v.point)*v.coeff,
	//↑注意！「,」で終わらせることによって、次の行もまだ文の続きですよ、と言える
   )
 }


 fmt.Fprintf(writer, "\n*** 構造体を埋め込んだ構造体(型の拡張) ***")
 //拡張した部分に値を追加
 vip_yumiko := vip{yumiko, 30}

 //ポイントを合計した変数を用意
 vip_point := vip_yumiko.member.point + vip_yumiko.vip_point

 fmt.Fprintf(writer, "\n%sさんはVIPなのでポイントは%d点",vip_yumiko.member.name, vip_point)

 vip_effective_point := float64(vip_point)*vip_yumiko.member.coeff
 fmt.Fprintf(writer, "\n%.2f",vip_effective_point)
 fmt.Fprintf(writer, "\n有効ポイントは%.2f点",vip_effective_point)
}

func main(){
	http.HandleFunc("/hello",hello)
	http.HandleFunc("/algebra",algebra)
	http.HandleFunc("/arrays",arrays)
	http.HandleFunc("/slices",slices)
	http.HandleFunc("/struct_members",struct_members)
	http.ListenAndServe(":8090",nil)
}