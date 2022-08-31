package crm

// CrmMember 结构体
type CrmMember struct {
	// 交易关闭的金额
	CloseTradeAmount string `json:"close_trade_amount,omitempty" xml:"close_trade_amount,omitempty"`
	// 显示会员的状态，normal正常，blacklist黑名单
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 会员拥有的所有分组
	GroupIds string `json:"group_ids,omitempty" xml:"group_ids,omitempty"`
	// 最后交易时间
	LastTradeTime string `json:"last_trade_time,omitempty" xml:"last_trade_time,omitempty"`
	// 城市.请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&amp;amp;path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// openuid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 交易成功的金额
	TradeAmount string `json:"trade_amount,omitempty" xml:"trade_amount,omitempty"`
	// 平均客单价.
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 等级名称
	GradeName string `json:"grade_name,omitempty" xml:"grade_name,omitempty"`
	// 购买的宝贝件数
	ItemNum int64 `json:"item_num,omitempty" xml:"item_num,omitempty"`
	// 交易关闭的宝贝件数
	ItemCloseCount int64 `json:"item_close_count,omitempty" xml:"item_close_count,omitempty"`
	// 关系来源，1交易成功，2未成交，3卖家主动吸纳
	RelationSource int64 `json:"relation_source,omitempty" xml:"relation_source,omitempty"`
	// 交易关闭的的笔数
	CloseTradeCount int64 `json:"close_trade_count,omitempty" xml:"close_trade_count,omitempty"`
	// 最后一次交易的订单号.注:该字段从2014.4.23之后不再返回.
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35.注:请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&amp;amp;path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 会员等级，0未非会员，其余对应等级名称取grade_name
	Grade int64 `json:"grade,omitempty" xml:"grade,omitempty"`
	// 交易成功笔数
	TradeCount int64 `json:"trade_count,omitempty" xml:"trade_count,omitempty"`
}
