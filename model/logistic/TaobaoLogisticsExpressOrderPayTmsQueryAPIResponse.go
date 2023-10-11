package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse 上门取退运费支付状态查询接口 API返回值
// taobao.logistics.express.order.pay.tms.query
//
// 上门取退运费支付状态查询接口
type TaobaoLogisticsExpressOrderPayTmsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel
}

// TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel is 上门取退运费支付状态查询接口 成功返回结果
type TaobaoLogisticsExpressOrderPayTmsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_order_pay_tms_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 业务处理结果
	Data *Tms2MscPayQueryResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
	// 是	系统成功失败 true|false
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}
