package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtPayMerchantFundstypeModifyAPIResponse 修改摊位分账类型 API返回值
// tmall.nrt.pay.merchant.fundstype.modify
//
// 修改摊位分账类型
type TmallNrtPayMerchantFundstypeModifyAPIResponse struct {
	model.CommonResponse
	TmallNrtPayMerchantFundstypeModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtPayMerchantFundstypeModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtPayMerchantFundstypeModifyAPIResponseModel).Reset()
}

// TmallNrtPayMerchantFundstypeModifyAPIResponseModel is 修改摊位分账类型 成功返回结果
type TmallNrtPayMerchantFundstypeModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_pay_merchant_fundstype_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统参数
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtPayMerchantFundstypeModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtPayMerchantFundstypeModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtPayMerchantFundstypeModifyAPIResponse)
	},
}

// GetTmallNrtPayMerchantFundstypeModifyAPIResponse 从 sync.Pool 获取 TmallNrtPayMerchantFundstypeModifyAPIResponse
func GetTmallNrtPayMerchantFundstypeModifyAPIResponse() *TmallNrtPayMerchantFundstypeModifyAPIResponse {
	return poolTmallNrtPayMerchantFundstypeModifyAPIResponse.Get().(*TmallNrtPayMerchantFundstypeModifyAPIResponse)
}

// ReleaseTmallNrtPayMerchantFundstypeModifyAPIResponse 将 TmallNrtPayMerchantFundstypeModifyAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtPayMerchantFundstypeModifyAPIResponse(v *TmallNrtPayMerchantFundstypeModifyAPIResponse) {
	v.Reset()
	poolTmallNrtPayMerchantFundstypeModifyAPIResponse.Put(v)
}
