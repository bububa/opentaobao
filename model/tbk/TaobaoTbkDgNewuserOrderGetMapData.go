package tbk

// TaobaoTbkDgNewuserOrderGetMapData 结构体
type TaobaoTbkDgNewuserOrderGetMapData struct {
	// 复购订单，仅适用于手淘拉新
	Orders []OrderData `json:"orders,omitempty" xml:"orders>order_data,omitempty"`
	// 新注册时间，仅淘宝拉新适用
	RegisterTime string `json:"register_time,omitempty" xml:"register_time,omitempty"`
	// 当前活动为淘宝拉新活动时，bind_time为新激活时间； 当前活动为支付宝拉新活动时，bind_time为绑定时间。
	BindTime string `json:"bind_time,omitempty" xml:"bind_time,omitempty"`
	// 首购时间，仅淘宝，天猫拉新适用
	BuyTime string `json:"buy_time,omitempty" xml:"buy_time,omitempty"`
	// 新人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 分享用户(unionid)，仅淘宝，天猫拉新适用
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// 来源媒体名称
	MemberNick string `json:"member_nick,omitempty" xml:"member_nick,omitempty"`
	// 来源站点名称
	SiteName string `json:"site_name,omitempty" xml:"site_name,omitempty"`
	// 来源广告位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 确认收货时间，仅天猫拉新适用
	AcceptTime string `json:"accept_time,omitempty" xml:"accept_time,omitempty"`
	// 领取红包时间，仅天猫拉新适用
	ReceiveTime string `json:"receive_time,omitempty" xml:"receive_time,omitempty"`
	// 拉新成功时间，仅支付宝拉新适用
	SuccessTime string `json:"success_time,omitempty" xml:"success_time,omitempty"`
	// 活动类型，taobao-淘宝 alipay-支付宝 tmall-天猫
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 日期，格式为"20180202"
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 绑卡日期，仅适用于手淘拉新
	BindCardTime string `json:"bind_card_time,omitempty" xml:"bind_card_time,omitempty"`
	// loginTime
	LoginTime string `json:"login_time,omitempty" xml:"login_time,omitempty"`
	// 使用权益时间
	UseRightsTime string `json:"use_rights_time,omitempty" xml:"use_rights_time,omitempty"`
	// 领取权益时间
	GetRightsTime string `json:"get_rights_time,omitempty" xml:"get_rights_time,omitempty"`
	// 渠道关系id
	RelationId string `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 新人状态， 当前活动为淘宝拉新活动时，1: 新注册，2:激活，3:首购，4:确认收货； 当前活动为支付宝实名活动时，1：已绑定，2：拉新成功，3：无效用户；当前活动为支付宝新登活动时，3：手淘首购，4：手淘确认收货；当前活动为天猫拉新活动时，2:已领取，3:已首购，4:已收货
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 订单淘客类型:1.淘客订单；2.非淘客订单，仅淘宝，天猫拉新适用
	OrderTkType int64 `json:"order_tk_type,omitempty" xml:"order_tk_type,omitempty"`
	// 来源媒体ID(pid中mm_1_2_3)中第1位
	MemberId int64 `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// 来源站点ID(pid中mm_1_2_3)中第2位
	SiteId int64 `json:"site_id,omitempty" xml:"site_id,omitempty"`
	// 来源广告位ID(pid中mm_1_2_3)中第3位
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 淘宝订单id，仅淘宝，天猫拉新适用
	TbTradeParentId int64 `json:"tb_trade_parent_id,omitempty" xml:"tb_trade_parent_id,omitempty"`
	// 银行卡是否是绑定状态：1-绑定，0-未绑定
	IsCardSave int64 `json:"is_card_save,omitempty" xml:"is_card_save,omitempty"`
}
