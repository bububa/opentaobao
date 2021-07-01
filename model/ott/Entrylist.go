package ott

// Entrylist 结构体
type Entrylist struct {
	// 入口ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 入口名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 入口图标
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 跳转行为
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 行为扩展
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
}
