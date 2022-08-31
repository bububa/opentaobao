package alihouse

// EtcThemeDetailExcelDto 结构体
type EtcThemeDetailExcelDto struct {
	// 业务id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部小区id
	OuterProjectId string `json:"outer_project_id,omitempty" xml:"outer_project_id,omitempty"`
	// 类型(1:楼盘 2:房源 3:文章 4:视频 5:评测 6:banner 7:品牌 8:抽奖 9:二手房小区 10:其他 11:租房 12:内容 13: &#34;交易商品&#34; 14: 二手房委托云主题 15: 租房委托云主题
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 10-集中式模式 11-分散式模式
	ProjectType int64 `json:"project_type,omitempty" xml:"project_type,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}
