package feedflow

// AdgroupQueryDto 
type AdgroupQueryDto struct {
    // 单元id列表
    AdgroupIdList   []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
    // 计划id列表
    CampaignIdList   []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
    // 分页起始位置
    Offset   int64 `json:"offset,omitempty" xml:"offset,omitempty"`
    // 每页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // PAUSE("投放暂停"),START("投放开始"),ERMINATE("投放停止")
    StatusList   []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
    // 单元名称
    AdgroupName   string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
}
