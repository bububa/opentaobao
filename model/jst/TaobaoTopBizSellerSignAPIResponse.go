package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopBizSellerSignAPIResponse 淘宝订单履约-商家erp签约 API返回值
// taobao.top.biz.seller.sign
//
// 淘宝订单履约-商家erp签约，包含各场景的签约
type TaobaoTopBizSellerSignAPIResponse struct {
	model.CommonResponse
	TaobaoTopBizSellerSignAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopBizSellerSignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopBizSellerSignAPIResponseModel).Reset()
}

// TaobaoTopBizSellerSignAPIResponseModel is 淘宝订单履约-商家erp签约 成功返回结果
type TaobaoTopBizSellerSignAPIResponseModel struct {
	XMLName xml.Name `xml:"top_biz_seller_sign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败code
	CallErrCode string `json:"call_err_code,omitempty" xml:"call_err_code,omitempty"`
	// 失败msg
	CallErrMsg string `json:"call_err_msg,omitempty" xml:"call_err_msg,omitempty"`
	// 调用结果
	CallResult bool `json:"call_result,omitempty" xml:"call_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopBizSellerSignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CallErrCode = ""
	m.CallErrMsg = ""
	m.CallResult = false
}

var poolTaobaoTopBizSellerSignAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopBizSellerSignAPIResponse)
	},
}

// GetTaobaoTopBizSellerSignAPIResponse 从 sync.Pool 获取 TaobaoTopBizSellerSignAPIResponse
func GetTaobaoTopBizSellerSignAPIResponse() *TaobaoTopBizSellerSignAPIResponse {
	return poolTaobaoTopBizSellerSignAPIResponse.Get().(*TaobaoTopBizSellerSignAPIResponse)
}

// ReleaseTaobaoTopBizSellerSignAPIResponse 将 TaobaoTopBizSellerSignAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopBizSellerSignAPIResponse(v *TaobaoTopBizSellerSignAPIResponse) {
	v.Reset()
	poolTaobaoTopBizSellerSignAPIResponse.Put(v)
}
