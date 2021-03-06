package menu

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func DaysInMonth(year,month int) int {
	switch time.Month(month) {
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) { // leap year
			return 29
		}
		return 28
	default:
		return 31
	}
}

func IsWorkDay(year,month,day int, excludeDays ...int)bool{
	t := time.Date(year+1911, time.Month(month),day,0,0,0,0,time.UTC).Weekday()
	weekday := int(t)
	for _,d := range excludeDays {
		if weekday == d {
			return false
		}
	}
	return true
}

func WorkDayList(year, month int,excludeDays ...int)[]string{
	buffer := DaysInMonth(year,month)
	daylist := []string{}
	for i:=1;i<=buffer;i++{
		if ok := IsWorkDay(year,month,i,excludeDays...);ok{
			day := strconv.Itoa(i)
			daylist = append(daylist,day)
		}
	}
	return daylist
}

func TimestringTransfertoInt(opts map[string]*MenuOptions)(map[string]interface{},error){
	r := make(map[string]interface{},len(opts))
	for k,v := range opts {
		if k == "ExcludeDays" {
			n,err := ExcludeDaystoInt(v)
			if err != nil {
				return nil,err
			}
			r[k] = n
			continue
		}
		n,err := strconv.Atoi(v.Value)
		if err != nil{
			return nil,err
		}
		r[k] = n
	}
	return r, nil
}

func ExcludeDaystoInt (optsValue *MenuOptions)([]int,error){
	r := []int{}
	d := strings.Split(optsValue.Value,",")
	for _,v := range d {
		if v == ""{
			continue
		}

		vr := strings.TrimSpace(v)
		n,err := strconv.Atoi(vr)
		if err != nil{
			return nil,err
		}
		r = append(r,n)
	}
	return r, nil
}






func outCalcTimefmtList (hourZone,leftHour int)([]string) {
	TimefmtList := []string{}
	for i:=0;i<=hourZone;i++{
		if i == hourZone && leftHour == 0 {
			break
		}
		t := ""
		switch state:=i%2;state{
		case 0:
			if i != hourZone {
				t += "8-12"
			} else {
				left := 8+leftHour
				t += fmt.Sprintf("8-%d",left)
				continue
			}
			i++
			if i == hourZone && leftHour == 0 {
				break
			}
			t+=","
			fallthrough
		case 1:
			if i != hourZone {
				t += "13-17"
			} else {
				left := 13+leftHour
				t += fmt.Sprintf("13-%d",left)
			}
		}
		TimefmtList = append(TimefmtList,t)
	}
	return TimefmtList
}