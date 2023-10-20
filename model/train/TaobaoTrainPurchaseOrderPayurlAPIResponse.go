package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainpurchaseorderpayurlAPIResponse 火车票采购商接口-获取支付链接 API返回值
// taobao.train.purchase.order.payurl
//
// 火车票采购商接口-获取支付链接
type TaobaotrainpurchaseorderpayurlAPIResponse struct {
	model.CommonResponse
	TaobaotrainpurchaseorderpayurlAPIResponseModel
}

// TaobaotrainpurchaseorderpayurlAPIResponseModel is 火车票采购商接口-获取支付链接 成功返回结果
type TaobaotrainpurchaseorderpayurlAPIResponseModel struct {
	XMLName xml.Name `xml:"train_purchase_order_payurl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 失败msg补充
	InnerMsg string `json:"inner_msg,omitempty" xml:"inner_msg,omitempty"`
	// 支付链接地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 失败msg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
