package fortuneTell

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type sliceSign []struct {
	Ji     string
	Qianci string
	Jie1   string
	Jie2   string
}
type Sign struct {
	Ji     string
	Qianci string
	Jie1   string
	Jie2   string
}

var fileJson = "signs.json"
var signs sliceSign
var isAsked map[string]int
var resetTime time.Time
var resetTimer  = time.NewTicker(1*time.Minute)

func init() {
	if !exist(fileJson) {
		err := DownloadJson(fileJson)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Open(fileJson)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(content, &signs)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().Unix())
	ResetData()
	SetResetTime(time.Date(0,00,0,23,30,0,0,time.Local))
	go doReset()
}

func ResetData(){
	isAsked = make(map[string]int,0)
}

func SetResetTime(reset time.Time) {
	resetTime = reset
}

func doReset(){
	for{
		select {
		case curr:=<-resetTimer.C:
			if curr.Hour() == resetTime.Hour() && curr.Minute() == resetTime.Minute() {
				ResetData()
				//fmt.Println("Do reset")
			}
			//fmt.Println(curr)
		}
	}
}

func AskSign(key string) (Sign, error) {
	defer func() { isAsked[key]++ }()
	if isAsked[key] == 0 {
		return getASign(), nil
	}
	return Sign{}, errors.New("has asked for sliceSign")
}

func IsAsked(key string) int {
	return IsAsked(key)
}

func getASign() Sign {
	var index = rand.Intn(len(signs))
	return signs[index]
}

func ParseJi(Ji string) string {
	var JiMap = map[string]int{
		"下下": 0,
		"中下": 1,
		"中平": 2,
		"中吉": 3,
		"上吉": 4,
		"上上": 5,
		"大吉": 6,
	}
	var count = JiMap[Ji]
	var ret string
	for i := 0; i < 6; i++ {
		var cell string
		if i < count {
			cell = "★"
		} else {
			cell = "☆"
		}
		ret += cell
	}
	return ret
}

func DownloadJson(saveTo string) error {
	var mSigns = make(sliceSign, 0)
	durl := "https://www.buyiju.com/guandi/%d.html"
	for i := 1; i <= 100; i++ {
		doc, err := goquery.NewDocument(fmt.Sprintf(durl, i))
		if err != nil {
			return err
		}
		var aSign Sign
		doc.Find("div.content").Find("p").Each(func(i int, selection *goquery.Selection) {
			tmpText := selection.Text()
			switch i {
			case 0:
				leftIndex := strings.Index(tmpText, "(")
				rightIndex := strings.Index(tmpText, ")")
				aSign.Ji = tmpText[leftIndex+1 : rightIndex]
			case 3:
				aSign.Qianci = tmpText
			case 7:
				aSign.Jie1 = tmpText
			case 9:
				left := strings.Index(tmpText, "：")
				aSign.Jie2 = tmpText[left+3:]
			}
		})
		mSigns = append(mSigns, aSign)
	}

	ret,_  := json.Marshal(mSigns)
	f, err := os.Create(saveTo)
	if err != nil {
		return err
	}
	f.Write(ret)
	return f.Close()
}

func exist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
