/*
 * MIT License
 *
 * Copyright (c) 2021 ashing
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"leetcode-question-today/api"
	"leetcode-question-today/msgpush"
)

var (
	wecom string // wecom 通知链接
	help  bool   // 帮助
	total int
)

func init() {
	flag.StringVar(&wecom, "wecom", "", "wecom webhook token")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.IntVar(&total, "total", 200, "total")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stdout, `leetcode-question-today - leetcode 每日一题推送
Usage: leetcode-question-today [-h help]
Options:
`)
	flag.PrintDefaults()
}

const msgTemplate = `每日一题(%s)
Title: %s
Difficulty: %s
AcRate: %.2f
Tags: %s
Link: %s
LinkCN: %s`

const TopTwo200 = `热门题目第 %s 道
Title: %s
Difficulty: %s
AcRate: %.2f
Tags: %s
Link: %s
LinkCN: %s`

func main() {
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}

	// 获取每日一题，如果有则推送即可
	resp, err := api.GetTodayQuestion(context.TODO())
	if err != nil {
		log.Printf("获取每日一题发生错误: %v\n", err)
		return
	}
	qs := []api.QuestionLightNode{}
	batchSize := 100
	for i := 0; i < total; i += batchSize {
		respAll, err := api.GetALLQuestions(i, min(i+batchSize, total))
		if err != nil {
			log.Printf("api.GetALLQuestionsV2 %v\n", err)
			return
		}
		qs = append(qs, respAll.P.Questions...)
	}
	filter := filterPaid(qs)
	if len(resp.TodayRecord) <= 0 {
		log.Printf("todayRecord 长度为 0,请检查\n")
		return
	}
	if len(filter) <= 0 {
		log.Printf("filter 长度为 0,请检查\n")
		return
	}
	fmt.Println(filter)

	today := resp.TodayRecord[0]
	date := today.Date
	for i := 0; i < 3; i++ {
		index := rand.New(rand.NewSource(time.Now().UnixNano())).Int() % len(filter)
		q := filter[index]
		diff := q.Difficulty
		title := q.TitleCn
		link := fmt.Sprintf("%s/problems/%s", api.Leetcode, q.TitleSlug)
		linkCn := fmt.Sprintf("%s/problems/%s", api.LeetcodeCn, q.TitleSlug)
		tags := make([]string, 0)
		for _, tag := range q.TopicTags {
			tags = append(tags, fmt.Sprintf("%s(%s)", tag.NameTranslated, tag.Name))
		}
		tagsValue := strings.Join(tags, "、")
		content := fmt.Sprintf(TopTwo200, q.FrontendQuestionId, title, diff, q.AcRate, tagsValue, link, linkCn)

		log.Println(content)

		if wecom != "" {
			w := msgpush.NewWeCom(wecom)
			_ = w.SendText(content, []string{"hongyu.an", "quan.ren", "binyuan.rong", "ming.wen"})
		}
	}

	difficulty := today.Question.Difficulty
	title := fmt.Sprintf("%s(%s)", today.Question.TitleCn, today.Question.Title)
	tags := make([]string, 0)
	for _, tag := range today.Question.TopicTags {
		tags = append(tags, fmt.Sprintf("%s(%s)", tag.NameTranslated, tag.Name))
	}
	tagsValue := strings.Join(tags, "、")
	link := fmt.Sprintf("%s/problems/%s", api.Leetcode, today.Question.TitleSlug)
	linkCn := fmt.Sprintf("%s/problems/%s", api.LeetcodeCn, today.Question.TitleSlug)

	content := fmt.Sprintf(msgTemplate, date, title, difficulty, today.Question.AcRate, tagsValue, link, linkCn)

	log.Println(content)

	if wecom != "" {
		w := msgpush.NewWeCom(wecom)
		_ = w.SendText(content, []string{"bin.zhang"})
	}

	return
}

func filterPaid(in []api.QuestionLightNode) (out []api.QuestionLightNode) {
	out = make([]api.QuestionLightNode, 0)
	for _, q := range in {
		if q.PaidOnly {
			continue
		}
		out = append(out, q)
	}

	return
}
