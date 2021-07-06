package viapi

// Face 结构体
type Face struct {
	// 人脸Id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 相似人物的名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 相似概率
	Rate int64 `json:"rate,omitempty" xml:"rate,omitempty"`
}
