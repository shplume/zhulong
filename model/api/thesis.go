package api

import "time"

type CreateThesis struct {
	ChineseTitle  string   `json:"chineseTitle" binding:"required"`  // 论文中文标题
	EnglishTitle  string   `json:"englishTitle" binding:"required"`  // 论文英文标题
	Authors       []string `json:"authors" binding:"required"`       // 论文作者
	Teachers      []string `json:"teachers" binding:"required"`      // 论文指导教师
	FirstAdvance  string   `json:"firstAdvance" binding:"required"`  // 第一创新点
	SecondAdvance string   `json:"secondAdvance" binding:"required"` // 第二创新点
	ThirdAdvance  string   `json:"thirdAdvance" binding:"required"`  // 第三创新点
	Drawback      string   `json:"drawback" binding:"required"`      // 论文不足
}

type UploadThesis struct {
	ThesisId int `form:"thesisId" binding:"required"` // 论文 ID
}

type DownloadThesis struct {
	ThesisId int `json:"thesisId" binding:"required"` // 论文 ID
}

type ToBeReviewedThesisList struct {
	ID            int       `json:"id"`             // 论文 ID
	FileName      string    `json:"file_name"`      // 论文的文件名称
	FileState     int       `json:"file_state"`     // 论文状态
	UploadTime    time.Time `json:"upload_time"`    // 论文上传时间
	ChineseTitle  string    `json:"chinese_title"`  // 论文中文标题
	EnglishTitle  string    `json:"english_title"`  // 论文英文标题
	Authors       string    `json:"authors"`        // 论文作者
	Teachers      string    `json:"teachers"`       // 论文指导老师
	FirstAdvance  string    `json:"first_advance"`  // 论文第一创新点
	SecondAdvance string    `json:"second_advance"` // 论文第二创新点
	ThirdAdvance  string    `json:"third_advance"`  // 论文第三创新点
	Drawback      string    `json:"drawback"`       // 论文不足
}

type AllocationThesis struct {
	ThesisID  []int `json:"thesisIds" binding:"required"` // 论文 ID
	TeacherID int   `json:"teacherId" binding:"required"` // 老师 ID
}

type RandomAllocationThesis struct {
	ThesisID []int `json:"thesisIds" binding:"required"` // 论文 ID
}
