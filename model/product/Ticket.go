package product

// Ticket 结构体
type Ticket struct {
	// 审核原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 关于审核原因，更详细的说明
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 产品规格申请时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 产品规格审核单最后修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 产品规格ID
	SpecId int64 `json:"spec_id,omitempty" xml:"spec_id,omitempty"`
	// 创建人ID
	CreateUserId int64 `json:"create_user_id,omitempty" xml:"create_user_id,omitempty"`
	// 如果产品规格，需要商家审核，为商家审核用户ID
	AuditSellerId int64 `json:"audit_seller_id,omitempty" xml:"audit_seller_id,omitempty"`
	// 1, &#34;商家确认&#34;&lt;br/&gt;2, &#34;商家拒绝&#34;&lt;br/&gt;3, &#34;小二确认&#34;&lt;br/&gt;4, &#34;小二拒绝&#34;&lt;br/&gt;5, &#34;待商家处理&#34;&lt;br/&gt;6, &#34;商家审核超时&#34;&lt;br/&gt;7, &#34;待小二审核&#34;&lt;br/&gt;9, &#34;品牌商确认&#34;&lt;br/&gt;10, &#34;免审通过&#34;&lt;br/&gt;14, &#34;免审拒绝&#34;
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
