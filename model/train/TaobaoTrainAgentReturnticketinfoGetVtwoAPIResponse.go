package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponse 代理商获取退票详情回调 API返回值
// taobao.train.agent.returnticketinfo.get.vtwo
//
// 代理商获取退票详情回调
type TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponseModel
}

// TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponseModel is 代理商获取退票详情回调 成功返回结果
type TaobaoTrainAgentReturnticketinfoGetVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_returnticketinfo_get_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 证件照图片列表
	CertificateUrlList []string `json:"certificate_url_list,omitempty" xml:"certificate_url_list>string,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 退票处理截止时间
	RefundDeadline string `json:"refund_deadline,omitempty" xml:"refund_deadline,omitempty"`
	// 预留字段
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 退票类型：10：跑腿退
	OfflineRefundType int64 `json:"offline_refund_type,omitempty" xml:"offline_refund_type,omitempty"`
	// 淘宝子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 淘宝主订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 证件类型
	CertificateType int64 `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
	// 线下票跑腿费
	VipErrandReturnPrice int64 `json:"vip_errand_return_price,omitempty" xml:"vip_errand_return_price,omitempty"`
	// 是否线下票退票
	OfflineReturn bool `json:"offline_return,omitempty" xml:"offline_return,omitempty"`
	// 查询返回状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
