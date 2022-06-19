package define

import (
	"os"
	"time"
)

var (
	DefaultPage = "1"
	DefaultSize = "20"
)

var MailPassword = os.Getenv("MailPassword")
var MysqlDNS = os.Getenv("MysqlDNS")

type ProblemBasic struct {
	Identity          string      `json:"identity"`           // 问题表的唯一标识
	Title             string      `json:"title"`              // 问题标题
	Content           string      `json:"content"`            // 问题内容
	ProblemCategories []int       `json:"problem_categories"` // 关联问题分类表
	MaxRuntime        int         `json:"max_runtime"`        // 最大运行时长
	MaxMem            int         `json:"max_mem"`            // 最大运行内存
	TestCases         []*TestCase `json:"test_cases"`         // 关联测试用例表
}

type TestCase struct {
	Input  string `json:"input"`  // 输入
	Output string `json:"output"` // 输出
}

type ProblemCollect struct {
	Identity  string    `json:"identity"` // 问题表的唯一标识
	Title     string    `json:"title"`    // 问题标题
	CreatedAt time.Time `json:"created_at"`
	PassNum   int64     `json:"pass_num"`   // 通过次数
	SubmitNum int64     `json:"submit_num"` // 提交次数
}

type ExplanationList struct {
	Identity  string    `json:"identity"` // 问题表的唯一标识
	Title     string    `json:"title"`    // 问题标题
	CreatedAt time.Time `json:"created_at"`
	ReadNum   int64     `json:"read_num"`
	LikeNum   int64     `json:"like_num"`
}

type ExplanationCreate struct {
	ProblemIdentity string `json:"problemIdentity"`
	Title           string `json:"title"`
	Content         string `json:"content"`
}

type ExplanationModify struct {
	Identity     string `json:"identity"`
	UserIdentity string `json:"userIdentity"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}

type CommentCreate struct {
	ContentIdentity string `json:"contentIdentity"`
	Comment         string `json:"comment"`
}

type CommentList struct {
	Identity     string    `json:"identity"`
	UserIdentity string    `json:"userIdentity"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
}

type ReplyCreate struct {
	CommentIdentity string `json:"CommentIdentity"`
	Reply           string `json:"Reply"`
}

type ReplyList struct {
	Identity     string    `json:"identity"`
	UserIdentity string    `json:"userIdentity"`
	Reply        string    `json:"Reply"`
	CreatedAt    time.Time `json:"created_at"`
}
