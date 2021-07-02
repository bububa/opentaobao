package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderGetVtwoAPIResponse 代理商获取订单信息回调APIv2--增加鉴权校验 API返回值
// taobao.train.agent.order.get.vtwo
//
// 代理商获取订单信息回调API
type TaobaoTrainAgentOrderGetVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderGetVtwoAPIResponseModel
}

// TaobaoTrainAgentOrderGetVtwoAPIResponseModel is 代理商获取订单信息回调APIv2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentOrderGetVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_get_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ttp_order_id
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
	// 返回错误。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 火车票信息。
	Tickets []ToAgentTicketInfo `json:"tickets,omitempty" xml:"tickets>to_agent_ticket_info,omitempty"`
	// 整个订单的总价,包括每张票价及保险价格,价格精确到分,例如100元,输出为10000.
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 移动电话
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 是否需要保险邮件地址
	Mailing bool `json:"mailing,omitempty" xml:"mailing,omitempty"`
	// 保险邮件地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 联系人姓名
	RelationName string `json:"relation_name,omitempty" xml:"relation_name,omitempty"`
	// 如果是公司发票，需要公司名称，如果不需要公司名称，返回no
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 1-已付款，2-关闭，3-成功
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 最晚出票时间
	LatestIssueTime string `json:"latest_issue_time,omitempty" xml:"latest_issue_time,omitempty"`
	// 订单类型0:默认订单类型走代理商账号；1：走12306客户绑定的账号；2：线下邮寄票
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 纸质票类型: 1 靠窗,2 连坐,3 上铺,4 中铺,5 下铺,6 是否同包厢
	PaperType int64 `json:"paper_type,omitempty" xml:"paper_type,omitempty"`
	// 当下铺/靠窗/连坐无票时,是否支持非下铺/非靠窗/非连坐(0不接受,1接受)
	PaperBackup int64 `json:"paper_backup,omitempty" xml:"paper_backup,omitempty"`
	// 至少接受下铺/靠窗/连坐数量
	PaperLowSeatCount int64 `json:"paper_low_seat_count,omitempty" xml:"paper_low_seat_count,omitempty"`
	// 线下票收件人姓名
	TransportName string `json:"transport_name,omitempty" xml:"transport_name,omitempty"`
	// 线下票收件人手机号
	TransportPhone string `json:"transport_phone,omitempty" xml:"transport_phone,omitempty"`
	// 线下票收件人地址
	TransportAddress string `json:"transport_address,omitempty" xml:"transport_address,omitempty"`
	// 快递费（分）
	TransportPrice int64 `json:"transport_price,omitempty" xml:"transport_price,omitempty"`
	// 手续费总价（分）
	ServicePrice int64 `json:"service_price,omitempty" xml:"service_price,omitempty"`
	// 扩展字段
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 1A2B
	OnlineBookSeat string `json:"online_book_seat,omitempty" xml:"online_book_seat,omitempty"`
	// interchangeStation
	InterchangeStation string `json:"interchange_station,omitempty" xml:"interchange_station,omitempty"`
	// isMultiTrip
	IsMultiTrip string `json:"is_multi_trip,omitempty" xml:"is_multi_trip,omitempty"`
	// 是否需要发票
	NeedReceipt bool `json:"need_receipt,omitempty" xml:"need_receipt,omitempty"`
	// 是否接受非定制坐席 1:是 0:否
	AcceptNoVipCustom int64 `json:"accept_no_vip_custom,omitempty" xml:"accept_no_vip_custom,omitempty"`
	// 用户可接受的最少定制票数量
	VipCustomMinSeatCount int64 `json:"vip_custom_min_seat_count,omitempty" xml:"vip_custom_min_seat_count,omitempty"`
	// 定制票类型 1:下铺 2:下铺or中铺 3:过道 4:靠窗
	VipCustomType int64 `json:"vip_custom_type,omitempty" xml:"vip_custom_type,omitempty"`
}
