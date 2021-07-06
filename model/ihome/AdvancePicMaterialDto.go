package ihome

// AdvancePicMaterialDto 结构体
type AdvancePicMaterialDto struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 图片oss地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 方案id
	CaseId string `json:"case_id,omitempty" xml:"case_id,omitempty"`
	// 宽
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 高
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 企业id
	EntId int64 `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 0 普通图 1 全景图 2 鸟瞰图
	Flag int64 `json:"flag,omitempty" xml:"flag,omitempty"`
}
