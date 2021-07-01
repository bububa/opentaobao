package baichuanctg

// AlibabaBaichuanCtgToutiaoContentData 结构体
type AlibabaBaichuanCtgToutiaoContentData struct {
	// 内容标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 内容发布时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 内容来源：头条
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 内容概要描述
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 内容封面图
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 内容链接（可能为空）
	ContentUrl string `json:"content_url,omitempty" xml:"content_url,omitempty"`
	// 内容来源说明
	OrgSource string `json:"org_source,omitempty" xml:"org_source,omitempty"`
	// 达人nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 微博大V账户id
	WbUid string `json:"wb_uid,omitempty" xml:"wb_uid,omitempty"`
	// 内容创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 返回的埋点参数
	Ybhpss string `json:"ybhpss,omitempty" xml:"ybhpss,omitempty"`
	// 内容正文信息
	Body string `json:"body,omitempty" xml:"body,omitempty"`
	// 内容修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 返回商品列表信息
	ItemUrlList string `json:"item_url_list,omitempty" xml:"item_url_list,omitempty"`
	// 内容id
	ContentId string `json:"content_id,omitempty" xml:"content_id,omitempty"`
	// 淘宝达人id
	TbUid string `json:"tb_uid,omitempty" xml:"tb_uid,omitempty"`
}
