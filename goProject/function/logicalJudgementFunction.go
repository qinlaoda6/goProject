package function

import "fmt"

// LeapYearJudge
// @description 闰年判断
// @param year
func LeapYearJudge(year int) {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println(year, "是闰年！")
	} else {
		fmt.Println(year, "不是闰年！")
	}
}

// SwitchJudge
// @description switch 判断使用
// @param grade
func SwitchJudge(grade float32) {
	switch {
	case grade > 90:
		fmt.Println("成绩优秀！")
	case grade >= 70:
		fmt.Println("成绩优良")
	case grade >= 60:
		fmt.Println("成绩及格")
	case grade == 59, grade == 58, grade == 57:
		fmt.Println("还差一点点,加油哦")
	default:
		fmt.Println("不及格！")
	}
}

// SingularJudge
// @description 计算奇数
// @param num
func SingularJudge(num int) {
	var num1 int
	fmt.Print("奇数数是:")
	for i := 0; i < num; i++ {
		if i%2 == 0 {
			num1++
			continue //结束本次循环
		}
		fmt.Print(" ", i)
	}
	fmt.Printf("\n奇数总数有%v个\n", num1)
}
