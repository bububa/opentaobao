package xiamicontent

// CollectDto 结构体
type CollectDto struct {
	// 歌单名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 主题
	Theme string `json:"theme,omitempty" xml:"theme,omitempty"`
	// 场景
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 歌单id
	CollectId int64 `json:"collect_id,omitempty" xml:"collect_id,omitempty"`
	// 业务方id
	PartnerId int64 `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
}
