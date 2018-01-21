package leetcodegraphql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// BaseQuestion the structure of the base question
type BaseQuestion struct {
	Problems          Problems `json:"-"`
	Referer           string   `json:"-"`
	Codes             Codes    `json:"-"`
	QuestionID        string   `json:"questionId"`
	QuestionTitle     string   `json:"questionTitle"`
	Content           string   `json:"content"`
	Difficulty        string   `json:"difficulty"`
	DiscussURL        string   `json:"discussUrl"`
	CategoryTitle     string   `json:"categoryTitle"`
	SubmitURL         string   `json:"submitUrl"`
	InterpretURL      string   `json:"interpretUrl"`
	CodeDefinition    string   `json:"codeDefinition"`
	MetaData          string   `json:"metaData"`
	EnvInfo           string   `json:"envInfo"`
	Article           string   `json:"article"`
	QuestionDetailURL string   `json:"questionDetailUrl"`
	DiscussCategoryID string   `json:"discussCategoryId"`
}

// Valid returns true if valid question
func (q BaseQuestion) Valid() bool {
	return q.QuestionID != "" && q.QuestionTitle != ""
}

// GetCodeDefinition returns code definition of question
func (q BaseQuestion) GetCodeDefinition(lang string) (code string, err error) {
	if c := q.Codes.Code(lang); c != nil {
		code = c.DefaultCode
		return
	}
	return
}

// GetEnvInfo returns env info
func (q BaseQuestion) GetEnvInfo(lang string) (info []string, err error) {
	s, err := strconv.Unquote(strconv.Quote(q.EnvInfo))
	if err != nil {
		return
	}

	m := make(map[string][]string)
	if err = json.Unmarshal([]byte(s), &m); err != nil {
		return
	}

	if temp, ok := m[lang]; ok && len(temp) != 0 {
		info = temp
	}
	return
}

// Do do requesting and parse the response data
func (q *BaseQuestion) Do(key string) error {
	titleSlug := key

	// try to parse id
	if q.Problems.Do() == nil {
		if s := q.Problems.StatStatus(key); s != nil {
			titleSlug = s.Stat.QuestionTitleSlug
		}
	}

	// parse title slug
	body := strings.NewReader(`{"query":"query getQuestionDetail($titleSlug: String!) {\n  isCurrentUserAuthenticated\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    questionTitle\n    questionTitleSlug\n    content\n    difficulty\n    stats\n    contributors\n    companyTags\n    topicTags\n    similarQuestions\n    discussUrl\n    mysqlSchemas\n    randomQuestionUrl\n    sessionId\n    categoryTitle\n    submitUrl\n    interpretUrl\n    codeDefinition\n    sampleTestCase\n    enableTestMode\n    metaData\n    enableRunCode\n    enableSubmit\n    judgerAvailable\n    infoVerified\n    envInfo\n    urlManager\n    article\n    questionDetailUrl\n    discussCategoryId\n    discussSolutionCategoryId\n  }\n  subscribeUrl\n  isPremium\n  loginUrl\n}\n","variables":{"titleSlug":"` +
		titleSlug + `"},"operationName":"getQuestionDetail"}`)
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", body)
	if err != nil {
		return err
	}
	q.Referer = fmt.Sprintf(
		"https://leetcode.com/problems/%s/description/",
		titleSlug,
	)
	req.Header.Set("x-csrftoken", guestToken)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("referer", q.Referer)
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	req.AddCookie(&http.Cookie{
		Name:    "csrftoken",
		Value:   guestToken,
		Path:    "/",
		Domain:  ".leetcode.com",
		Secure:  true,
		Expires: time.Now(),
	})
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = json.Unmarshal(data, &Response{
		Data: ResponseData{
			Question: q,
		},
	}); err != nil {
		return err
	}

	s, err := strconv.Unquote(strconv.Quote(q.CodeDefinition))
	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(s), &q.Codes); err != nil {
		return err
	}

	return nil
}
