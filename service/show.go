package service

import (
	"context"
	"fmt"
	"github.com/iooikaak/frame/mq/rocketmq"
	"github.com/iooikaak/frame/snowflake"
	"github.com/iooikaak/frame/xlog"
	model "github.com/iooikaak/job/common/model/mysql/net_edu"
	"github.com/iooikaak/job/config"
	"github.com/iooikaak/job/dao/mysql/net_edu"
	mq "github.com/iooikaak/job/dao/rocketmq"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"strconv"
)

var number = 9261

func (s *Service) tick(number int) {
	url := "https://www.hnkaoshiwang.com/tiku/stkemu" + strconv.Itoa(number) + ".html"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "www.hnkaoshiwang.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("CharSet", "gb2312")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := ConvertGB2312ToStr(body)

	d := net_edu.New(config.Conf)
	db := d.Db.Context(context.Background())
	id, _ := snowflake.NewID()
	a := d.InsertNetEdu(db.Debug(), &model.NetEdu{
		ID:     int64(id),
		Number: int64(number),
		Body:   str,
	})
	xlog.Infof("---%v---", a)
}

func ConvertGB2312ToStr(gbkStr []byte) string {
	//如果是[]byte格式的字符串，可以使用Bytes方法
	b, err := simplifiedchinese.GBK.NewDecoder().Bytes(gbkStr)
	if err != nil {
		return ""
	}
	return string(b)
}

func (s *Service) Show() {
	//xlog.Debug("===== number is %d =====", number)
	//if number > 10000 {
	//os.Exit(0)
	//}
	//s.tick(number)
	//number++
}

func (s *Service) StartShow() {
	selector := rocketmq.ConsumerMessageSelector{}
	err := s.mq.Consumer.RegisterHandle(mq.ConsumerTopicTest, selector, func(msgs []*rocketmq.MessageExt, ctx *rocketmq.ConsumeConcurrentlyContext) (rocketmq.ConsumeResult, error) {
		s.Show()
		xlog.Debugf("===== msgs is %+v =====", msgs)
		xlog.Debugf("===== ctx is %+v =====", ctx)
		return rocketmq.ConsumeSuccess, nil
	})
	if err != nil {
		xlog.Errorf("StartShow RegisterHandle failed err: %v", err)
		return
	}
	err = s.mq.Consumer.Start()
	if err != nil {
		xlog.Errorf("StartShow Start failed err: %v", err)
	}
}

func (s *Service) StopShow() {
	err := s.mq.Consumer.Stop()
	if err != nil {
		xlog.Errorf("StopShow failed err: %v", err)
	}
}
