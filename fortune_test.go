package fortune_telling

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strings"
	"testing"
)

const (
	DownUrl = "https://www.buyiju.com/guandi/%d.html"
)

func init() {
	fmt.Printf("test_init")
}

func downloadFortuneData(saveTo string) error {
	var mSigns = make([]Telling, 0)
	for i := 1; i <= 100; i++ {
		doc, err := goquery.NewDocument(fmt.Sprintf(DownUrl, i))
		if err != nil {
			return err
		}
		var telling Telling
		doc.Find("div.content").Find("p").Each(func(i int, selection *goquery.Selection) {
			tmpText := selection.Text()
			switch i {
			case 0:
				leftIndex := strings.Index(tmpText, "(")
				rightIndex := strings.Index(tmpText, ")")
				telling.Level = tmpText[leftIndex+1 : rightIndex]
			case 3:
				telling.Content = tmpText
			case 7:
				telling.Detail1 = tmpText
			case 9:
				left := strings.Index(tmpText, "：")
				telling.Detail2 = tmpText[left+3:]
			}
		})
		mSigns = append(mSigns, telling)
		log.Printf("%d get: %v", i, telling)
	}

	ret, _ := json.Marshal(mSigns)
	f, err := os.Create(saveTo)
	if err != nil {
		return err
	}
	_, _ = f.Write(ret)
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

func TestDownload(t *testing.T) {
	if !exist(fileJson) {
		err := downloadFortuneData(fileJson)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestAsk(t *testing.T) {
	for i := 0; i < 10; i++ {
		tell, err := Ask(fmt.Sprintf("%v", i))
		if err != nil {
			t.Errorf("ask: %v", err)
		} else {
			t.Logf("%v: %v", i, tell)
		}
	}
}
