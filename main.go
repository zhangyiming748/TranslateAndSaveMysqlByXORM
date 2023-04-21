package main

import (
	"TranslateAndSaveMysqlByXORM/model"
	"TranslateAndSaveMysqlByXORM/util"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/zhangyiming748/goini"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"strings"
	"time"
)

// The is the struct for the data returned by Bing.
const (
	configPath = "./conf.ini"
)

func SetLog(level string) {
	var opt slog.HandlerOptions
	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}

	}
	file := "baiduTranslate.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	//defer logf.Close() //如果不关闭可能造成内存泄露
	mylog := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
	slog.SetDefault(mylog)
}

type RES struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

// Verify the endpoint URI and replace the token string with a valid subscription key.
func main() {
	conf := goini.SetConfig(configPath)
	l, err := conf.GetValue("log", "level")
	if err != nil {
		return
	}
	SetLog(l)
	var Query string
	if len(os.Args) < 2 {
		open, err := os.ReadFile("doc.txt")
		if err != nil {
			return
		}
		slog.Info("读文件", slog.String("文件内容", string(open)))
		Query = string(open)
	} else {
		Query = os.Args[1]
	}
	w := new(model.Word)
	w.Other = Query
	word, has, err := w.FindByWord()
	if has {
		slog.Info("读取缓存")
	}
	result, err := os.OpenFile("chinese.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		slog.Error("写入翻译后的文本失败")
	}
	defer result.Close()
	if err != nil {
		return
	} else if has {
		slog.Info("查询结果", slog.String("src", Query), slog.String("dst", word.ZhCn))
		result.WriteString(strings.Join([]string{word.ZhCn, "\n"}, ""))
		return
	}
	appid, err := conf.GetValue("api", "appid")
	key, err := conf.GetValue("api", "key")
	from, err := conf.GetValue("api", "from")
	if from == "" {
		from = "auto"
	}
	to, err := conf.GetValue("api", "to")
	if err != nil {
		return
	}
	salt := time.Now().Format("0102150405")
	sign := strings.Join([]string{appid, Query, salt, key}, "")
	slog.Debug("加密之前", slog.String("拼接的字符串", sign))
	signMd5 := getMd5(sign)
	slog.Debug("加密之后", slog.String("拼接的字符串", signMd5))
	m := map[string]string{
		"q":     Query,
		"from":  from,
		"to":    to,
		"appid": appid,
		"salt":  salt,
		"sign":  signMd5,
	}
	slog.Debug("请求的参数", slog.Any("param", m))
	get, err := util.HttpGet(nil, m, "http://api.fanyi.baidu.com/api/trans/vip/translate")
	if err != nil {
		return
	}
	var res RES
	err = json.Unmarshal(get, &res)
	if err != nil {
		return
	}
	w.Other = res.TransResult[0].Src
	w.ZhCn = res.TransResult[0].Dst
	w.Kind = res.From
	w.CreateOne()
	slog.Info("自动判断语言", slog.String("源语言", res.From), slog.String("目标语言", res.To))
	slog.Info("查询结果", slog.String("src", res.TransResult[0].Src), slog.String("dst", res.TransResult[0].Dst))
	result.WriteString(strings.Join([]string{res.TransResult[0].Dst, "\n"}, ""))
}
func getMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
