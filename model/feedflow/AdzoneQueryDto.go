package feedflow

// AdzoneQueryDto 结构体
type AdzoneQueryDto struct {
	// 广告位id列表
	AdzoneIdList []int64 `json:"adzone_id_list,omitempty" xml:"adzone_id_list>int64,omitempty"`
	// 广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
