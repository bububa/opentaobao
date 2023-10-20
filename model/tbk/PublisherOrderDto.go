package tbk

import (
	"sync"
)

// PublisherOrderDto 结构体
type PublisherOrderDto struct {
	// 服务费信息（字段已废弃）
	ServiceFeeDtoList []ServiceFeeDto `json:"service_fee_dto_list,omitempty" xml:"service_fee_dto_list>service_fee_dto,omitempty"`
	// 补贴金额明细节点。解释：各补贴类型的类型名称、补贴比率、补贴金额、单笔补贴上限、补贴分成比率的详细说明
	SubsidyInfoDtoList []SubsidyDetailDto `json:"subsidy_info_dto_list,omitempty" xml:"subsidy_info_dto_list>subsidy_detail_dto,omitempty"`
	// 淘宝付款时间。解释：订单在淘宝付款的时间
	TbPaidTime string `json:"tb_paid_time,omitempty" xml:"tb_paid_time,omitempty"`
	// 付款时间。解释：订单付款的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单付款时间  （预售订单完尾款支付后，付款时间会自动更新为尾款支付的时间）
	TkPaidTime string `json:"tk_paid_time,omitempty" xml:"tk_paid_time,omitempty"`
	// 结算金额。解释：买家确认收货的付款金额（不包含运费金额）
	PayPrice string `json:"pay_price,omitempty" xml:"pay_price,omitempty"`
	// 结算预估收入。解释：结算预付收入=结算预估佣金收入+结算预估补贴收入
	PubShareFee string `json:"pub_share_fee,omitempty" xml:"pub_share_fee,omitempty"`
	// 子订单号。解释：买家通过购物车购买的每个商品对应的订单编号，此订单编号并未在淘宝买家后台透出，若平台类型为淘宝、口碑、饿了么等，则订单编号即为淘宝子订单编号、口碑订单编号、饿了么订单编号，以此类推
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 结算时间。解释：订单确认收货后且商家完成佣金支付的时间
	TkEarningTime string `json:"tk_earning_time,omitempty" xml:"tk_earning_time,omitempty"`
	// 佣金分成比率。解释：从佣金中分得的收益比率（含平台技术服务费比率）
	PubShareRate string `json:"pub_share_rate,omitempty" xml:"pub_share_rate,omitempty"`
	// 补贴比率。解释：指该笔订单上各类型补贴的补贴比率总和。补贴比率=a补贴比率+b补贴比率+…。举例：天猫补贴1%、飞猪补贴1%等，则该数值显示为2%
	SubsidyRate string `json:"subsidy_rate,omitempty" xml:"subsidy_rate,omitempty"`
	// 佣金提成。解释：佣金提成=佣金比率*佣金分成比率。指实际获得的佣金收益比率（含平台技术服务费）
	TkTotalRate string `json:"tk_total_rate,omitempty" xml:"tk_total_rate,omitempty"`
	// 类目名称。解释：商品所属的一级类目名称
	ItemCategoryName string `json:"item_category_name,omitempty" xml:"item_category_name,omitempty"`
	// 掌柜旺旺
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 推广者赚取佣金后支付给阿里妈妈的技术服务费用的比率
	AlimamaRate string `json:"alimama_rate,omitempty" xml:"alimama_rate,omitempty"`
	// 补贴类型。解释：各补贴类型的类型名称、补贴比率、补贴金额、单笔补贴上限、补贴分成比率的详细说明，如有多个补贴会一并告知，举例：淘宝特价版（补贴比率：1%，补贴金额1.00元，单笔补贴上限金额1000.00元，补贴分成比率：100.00%）、飞猪补贴（补贴比率：1%，补贴金额1.00元，单笔补贴上限金额1000.00元，补贴分成比率：100.00%）
	SubsidyType string `json:"subsidy_type,omitempty" xml:"subsidy_type,omitempty"`
	// 商品图片
	ItemImg string `json:"item_img,omitempty" xml:"item_img,omitempty"`
	// 付款预估收入。解释：付款预付收入=付款预估佣金收入+付款预估补贴收入
	PubSharePreFee string `json:"pub_share_pre_fee,omitempty" xml:"pub_share_pre_fee,omitempty"`
	// 付款金额。解释：买家拍下并付款金额（不包含运费金额、不包含未支付尾款的预售订单金额）；当预售订单完尾款支付后，付款金额会自动更新为订单总金额
	AlipayTotalPrice string `json:"alipay_total_price,omitempty" xml:"alipay_total_price,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 媒体名称。解释：“推广管理-媒体备案管理”中的媒体名称
	SiteName string `json:"site_name,omitempty" xml:"site_name,omitempty"`
	// 补贴金额。解释：指该笔订单上各类补贴的补贴金额总和。补贴金额=a补贴金额+b补贴金额+…=结算金额*a补贴比率+结算金额*b补贴比率+…。若（结算金额*a补贴比率）＞补贴类型a单笔补贴上限，则a补贴金额=补贴类型a单笔补贴上限，b补贴金额等其他补贴金额同理
	SubsidyFee string `json:"subsidy_fee,omitempty" xml:"subsidy_fee,omitempty"`
	// 平台技术服务费。解释：指该笔订单推广者在需支付给淘宝联盟的所有技术服务费费用总和。目前包含以下两类： 1、技术服务费说明：通过淘宝联盟平台推广所需支付的基础技术服务费；2、渠道专项服务费说明：开通渠道管理权限的推广者进行推广需要支付给淘宝联盟的专项技术服务费用。           平台技术服务费=付款金额*佣金比率*平台技术服务费比率。注意：若订单已结算则会按结算金额计算。（特别说明：补贴从2023-03-10之后都不再收取补贴技术服务费了，故本字段公式中不再提及补贴相关)
	AlimamaShareFee string `json:"alimama_share_fee,omitempty" xml:"alimama_share_fee,omitempty"`
	// 订单编号。解释：买家在购买后台显示的订单编号，若平台类型为淘宝、饿了么等，则订单编号即为淘宝订单编号、饿了么订单编号，以此类推
	TradeParentId string `json:"trade_parent_id,omitempty" xml:"trade_parent_id,omitempty"`
	// 平台类型。解释：订单所属平台类型，包括天猫、淘宝、聚划算、口碑等
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 创建时间。解释：订单创建的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间
	TkCreateTime string `json:"tk_create_time,omitempty" xml:"tk_create_time,omitempty"`
	// 产品类型。解释：若订单属于特定的产品类型，会进行说明，举例：该订单属于“联盟超级活动”、“品牌精选推广”、“自主推广”等，若无说明仅为“--”则为普通类型订单
	FlowSource string `json:"flow_source,omitempty" xml:"flow_source,omitempty"`
	// 成交平台。解释：成交来自于PC或无线
	TerminalType string `json:"terminal_type,omitempty" xml:"terminal_type,omitempty"`
	// 点击时间。解释：通过推广链接达到商品、店铺详情页的点击时间
	ClickTime string `json:"click_time,omitempty" xml:"click_time,omitempty"`
	// 商品单价
	ItemPrice string `json:"item_price,omitempty" xml:"item_price,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 推广位名称。解释：“推广管理-推广位管理”中的推广位名称
	AdzoneName string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
	// 佣金比率。解释：商品的佣金比率（含平台技术服务费比率，但不包含平台专项服务费比率）（特别说明：若想知道属于您名下的整体佣金比率，则整体佣金比率=佣金比率+平台专项服务费比率，为了简化展示报表，不再增加整体佣金收入字段，有需要的淘宝客可自行加和，举例：若整体佣金比率为10%，平台专项服务费比率为3%，则此处商品佣金比率显示为7%，即10%-3%=7%）
	TotalCommissionRate string `json:"total_commission_rate,omitempty" xml:"total_commission_rate,omitempty"`
	// 商品链接
	ItemLink string `json:"item_link,omitempty" xml:"item_link,omitempty"`
	// 店铺名称
	SellerShopTitle string `json:"seller_shop_title,omitempty" xml:"seller_shop_title,omitempty"`
	// 收入比率。解释：收入比率=佣金比率+补贴比率
	IncomeRate string `json:"income_rate,omitempty" xml:"income_rate,omitempty"`
	// 佣金金额。解释：如果该订单还未结算，则佣金金额=付款金额*佣金比率；如果该订单已经结算，则佣金金额=结算金额*佣金比率
	TotalCommissionFee string `json:"total_commission_fee,omitempty" xml:"total_commission_fee,omitempty"`
	// 定金付款时间。解释：预售时期，预售订单支付定金的付款时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间
	TkDepositTime string `json:"tk_deposit_time,omitempty" xml:"tk_deposit_time,omitempty"`
	// 定金淘宝付款时间。解释：预售时期，预售订单在淘宝支付定金的付款时间
	TbDepositTime string `json:"tb_deposit_time,omitempty" xml:"tb_deposit_time,omitempty"`
	// 定金付款金额。解释：预售时期，预售订单支付的定金金额
	DepositPrice string `json:"deposit_price,omitempty" xml:"deposit_price,omitempty"`
	// 开发者调用api的appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 非电商淘系子订单号
	TpOrderId string `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 营销类型。解释：该字段中视订单情况有单个或多个类型告知。 举例：联盟超级活动-优品、特价版客户端锁粉、特价版客户端推广等
	MarketingType string `json:"marketing_type,omitempty" xml:"marketing_type,omitempty"`
	// 订单更新时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 管理member新商品ID-B段
	ExtraMktId string `json:"extra_mkt_id,omitempty" xml:"extra_mkt_id,omitempty"`
	// unid（不对外开放）
	Unid string `json:"unid,omitempty" xml:"unid,omitempty"`
	// 专用（不对外开放）
	TalentPid string `json:"talent_pid,omitempty" xml:"talent_pid,omitempty"`
	// 买家拍下金额（不包含运费金额）
	TbGmvTotalPrice string `json:"tb_gmv_total_price,omitempty" xml:"tb_gmv_total_price,omitempty"`
	// 流量通untts（默认无，限定开放）
	Untts string `json:"untts,omitempty" xml:"untts,omitempty"`
	// 推广者身份。解释：二方：佣金收益的第一归属者；三方：从其他淘宝客佣金中进行分成的推广者
	TkOrderRole int64 `json:"tk_order_role,omitempty" xml:"tk_order_role,omitempty"`
	// 推广位ID。解释：“推广管理-推广位管理”中的pid中的最后一段数字，如pid=mm_1_2_3中的“3”这段数字
	AdzoneId int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
	// 维权标签。解释：若该订单产生了维权退款，则会打上“维权单”的提示。0 含义为非维权、1 含义为维权订单
	RefundTag int64 `json:"refund_tag,omitempty" xml:"refund_tag,omitempty"`
	// 推广者的账号id
	PubId int64 `json:"pub_id,omitempty" xml:"pub_id,omitempty"`
	// 商品数量
	ItemNum int64 `json:"item_num,omitempty" xml:"item_num,omitempty"`
	// 订单状态。状态类型：已付款（指订单已付款，但还未确认收货）、已收货（指订单已确认收货，但商家佣金未支付）、已结算（指订单已确认收货，且商家佣金已支付成功）、已失效（指订单关闭/订单佣金小于0.01元，订单关闭主要有：①买家超时未付款；②买家付款前，买家/商家取消了订单；③订单付款后发起售中退款成功；④仅支付定金，到期未支付尾款的预售订单）。注意：当订单商家有实收货款（含交易成功和交易关闭订单），联盟以商家实收货款进行佣金结算(满返/买返订单等特殊订单除外)。出参数字代表：3:订单结算，12:订单付款，13:订单失效，14:订单成功
	TkStatus int64 `json:"tk_status,omitempty" xml:"tk_status,omitempty"`
	// 媒体ID。解释：“推广管理-媒体备案管理”中的媒体ID，也是pid中的第二段数字，如pid=mm_1_2_3中的“2”这段数字
	SiteId int64 `json:"site_id,omitempty" xml:"site_id,omitempty"`
	// 会员运营ID。解释：会员运营管理功能中的会员运营ID，若该订单来自于会员的购买，则会展示对应会员运营ID
	SpecialId int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
	// 渠道关系id。解释：渠道管理功能中的渠道关系ID，若该订单来自于渠道方的推广，则会展示对应渠道关系ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 付款预估佣金收入。解释：付款预估佣金收入=付款金额*佣金提成。以买家付款金额为基数，预估您可能获得的付款佣金收入，包含平台技术服务费金额（最终发钱时会减去平台技术服务费）。注意：因买家退款等原因，可能与结算预估佣金收入不一致。（特别说明：若想知道属于您名下的整体佣金收入，则整体付款佣金收入=付款预估佣金收入+平台专项服务费，为了简化展示报表，不再增加整体佣金收入字段，有需要的淘宝客可自行加和）
	PubSharePreFeeForCommission float64 `json:"pub_share_pre_fee_for_commission,omitempty" xml:"pub_share_pre_fee_for_commission,omitempty"`
	// 结算预估佣金收入。解释：结算预估佣金收入=付款金额*佣金提成。以买家确认收货的付款金额为基数，预估您可能获得的结算佣金收入，包含技术服务费金额（最终发钱时会减去平台技术服务费）。注意：因买家退款、您违规推广等原因，可能与您最终收入不一致，最终收入以月结后您实际收到的为准。（特别说明：若想知道属于您名下的整体佣金收入，则整体结算佣金收入=结算预估佣金收入+平台专项服务费，为了简化展示报表，不再增加整体佣金收入字段，有需要的淘宝客可自行加和）
	PubShareFeeForCommission float64 `json:"pub_share_fee_for_commission,omitempty" xml:"pub_share_fee_for_commission,omitempty"`
	// 补贴分成比率。解释：从补贴中分得的收益比率
	PubShareRateForSdy float64 `json:"pub_share_rate_for_sdy,omitempty" xml:"pub_share_rate_for_sdy,omitempty"`
	// 补贴提成。解释：补贴提成=补贴比率*补贴分成比率。指实际获得的补贴收益比率
	TkTotalRateForSdy float64 `json:"tk_total_rate_for_sdy,omitempty" xml:"tk_total_rate_for_sdy,omitempty"`
	// 付款预估补贴收入=（付款金额*a补贴比率+付款金额*b补贴比率+……）*补贴分成比率。如果付款金额*a补贴比率＞补贴类型a单笔订单补贴上限，则付款金额*a补贴比率的值取补贴类型a单笔订单补贴上限，b补贴金额等其他类型补贴金额同理。指以买家付款金额为基数，预估您可能获得的补贴收入。因买家退款等原因，可能与结算预估补贴收入不一致。
	PubSharePreFeeForSdy float64 `json:"pub_share_pre_fee_for_sdy,omitempty" xml:"pub_share_pre_fee_for_sdy,omitempty"`
	// 结算预估补贴收入。解释：以买家确认收货的付款金额为基数，预估您可能获得的结算补贴收入。结算预估补贴收入=（结算金额*a补贴比率+结算金额*b补贴比率+……）*补贴分成比率。如果结算金额*a补贴比率＞补贴类型a单笔订单补贴上限，则结算金额*a补贴比率的值取补贴类型a单笔订单补贴上限，b补贴金额等其他类型补贴金额同理。注意：因买家退款、您违规推广等原因，可能与您最终收入不一致，最终收入以月结后您实际收到的为准
	PubShareFeeForSdy float64 `json:"pub_share_fee_for_sdy,omitempty" xml:"pub_share_fee_for_sdy,omitempty"`
	// 平台技术服务费明细节点。解释：各项平台技术服务费类型的类型名称、扣费比率、扣费金额的详细说明
	AlimmShareInfoDto *AlimmShareInfoDto `json:"alimm_share_info_dto,omitempty" xml:"alimm_share_info_dto,omitempty"`
	// 平台专项服务费比率。解释：指该笔订单推广者在各种特殊场景下需支付给淘宝联盟的所有专项服务费比率总和。平台专项服务费比率=专项服务费比率a+专项服务费比率b+…
	PlatformSpecialServiceRate float64 `json:"platform_special_service_rate,omitempty" xml:"platform_special_service_rate,omitempty"`
	// 平台专项服务费。解释：指该笔订单推广者在各种特殊场景下需支付给淘宝联盟的所有专项服务费用总和。目前包含以下两类：1、内容专项服务费说明：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给淘宝联盟的专项技术服务费用；2、流量专项服务费说明：流量通产品合作场景专项技术服务费，推广者在流量通产品能力下，进行流量投放推广需要支付给淘宝联盟的专项技术服务费用。           平台专项服务费=付款金额*平台专项服务费比率。注意：若订单已结算则会按结算金额计算
	PlatformSpecialServiceFee float64 `json:"platform_special_service_fee,omitempty" xml:"platform_special_service_fee,omitempty"`
	// 平台专项服务费明细节点。解释：各项平台专项服务费类型的类型名称、扣费比率、扣费金额的详细说明
	PlatformSpecialShareInfoDto *PlatformSpecialShareInfoDto `json:"platform_special_share_info_dto,omitempty" xml:"platform_special_share_info_dto,omitempty"`
}

var poolPublisherOrderDto = sync.Pool{
	New: func() any {
		return new(PublisherOrderDto)
	},
}

// GetPublisherOrderDto() 从对象池中获取PublisherOrderDto
func GetPublisherOrderDto() *PublisherOrderDto {
	return poolPublisherOrderDto.Get().(*PublisherOrderDto)
}

// ReleasePublisherOrderDto 释放PublisherOrderDto
func ReleasePublisherOrderDto(v *PublisherOrderDto) {
	v.ServiceFeeDtoList = v.ServiceFeeDtoList[:0]
	v.SubsidyInfoDtoList = v.SubsidyInfoDtoList[:0]
	v.TbPaidTime = ""
	v.TkPaidTime = ""
	v.PayPrice = ""
	v.PubShareFee = ""
	v.TradeId = ""
	v.TkEarningTime = ""
	v.PubShareRate = ""
	v.SubsidyRate = ""
	v.TkTotalRate = ""
	v.ItemCategoryName = ""
	v.SellerNick = ""
	v.AlimamaRate = ""
	v.SubsidyType = ""
	v.ItemImg = ""
	v.PubSharePreFee = ""
	v.AlipayTotalPrice = ""
	v.ItemTitle = ""
	v.SiteName = ""
	v.SubsidyFee = ""
	v.AlimamaShareFee = ""
	v.TradeParentId = ""
	v.OrderType = ""
	v.TkCreateTime = ""
	v.FlowSource = ""
	v.TerminalType = ""
	v.ClickTime = ""
	v.ItemPrice = ""
	v.ItemId = ""
	v.AdzoneName = ""
	v.TotalCommissionRate = ""
	v.ItemLink = ""
	v.SellerShopTitle = ""
	v.IncomeRate = ""
	v.TotalCommissionFee = ""
	v.TkDepositTime = ""
	v.TbDepositTime = ""
	v.DepositPrice = ""
	v.AppKey = ""
	v.TpOrderId = ""
	v.MarketingType = ""
	v.ModifiedTime = ""
	v.ExtraMktId = ""
	v.Unid = ""
	v.TalentPid = ""
	v.TbGmvTotalPrice = ""
	v.Untts = ""
	v.TkOrderRole = 0
	v.AdzoneId = 0
	v.RefundTag = 0
	v.PubId = 0
	v.ItemNum = 0
	v.TkStatus = 0
	v.SiteId = 0
	v.SpecialId = 0
	v.RelationId = 0
	v.PubSharePreFeeForCommission = 0
	v.PubShareFeeForCommission = 0
	v.PubShareRateForSdy = 0
	v.TkTotalRateForSdy = 0
	v.PubSharePreFeeForSdy = 0
	v.PubShareFeeForSdy = 0
	v.AlimmShareInfoDto = nil
	v.PlatformSpecialServiceRate = 0
	v.PlatformSpecialServiceFee = 0
	v.PlatformSpecialShareInfoDto = nil
	poolPublisherOrderDto.Put(v)
}
