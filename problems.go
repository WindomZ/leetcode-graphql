package leetcodegraphql

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Problems the structure of all problems
type Problems struct {
	StatStatusPairs []ProblemStatStatus `json:"stat_status_pairs"`
	NumTotal        int                 `json:"num_total"`
}

// ProblemStatStatus the structure of a problem status
type ProblemStatStatus struct {
	Stat       ProblemStat `json:"stat"`
	Difficulty struct {
		Level int `json:"level"`
	} `json:"difficulty"`
}

// ProblemStat the structure of a problem stat
type ProblemStat struct {
	TotalAcs            int    `json:"total_acs"`
	QuestionTitle       string `json:"question__title"`
	IsNewQuestion       bool   `json:"is_new_question"`
	QuestionArticleSlug string `json:"question__article__slug"`
	TotalSubmitted      int    `json:"total_submitted"`
	FrontendQuestionID  int    `json:"frontend_question_id"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionArticleLive bool   `json:"question__article__live"`
	QuestionHide        bool   `json:"question__hide"`
	QuestionID          int    `json:"question_id"`
}

// Do do requesting and parse the response data
func (p *Problems) Do() error {
	req, err := http.NewRequest("GET",
		"https://leetcode.com/api/problems/all/", nil)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = json.Unmarshal(data, p); err != nil {
		return err
	}

	return nil
}

// StatStatus returns ProblemStatStatus with id or title string
func (p Problems) StatStatus(s string) *ProblemStatStatus {
	s = strings.ToLower(strings.TrimSpace(s))

	if id, err := strconv.Atoi(s); err == nil {
		for _, pair := range p.StatStatusPairs {
			if id == pair.Stat.QuestionID {
				return &pair
			}
		}
	}

	for _, pair := range p.StatStatusPairs {
		if s == strings.ToLower(pair.Stat.QuestionTitle) ||
			s == strings.ToLower(pair.Stat.QuestionTitleSlug) {
			return &pair
		}
	}
	return nil
}
