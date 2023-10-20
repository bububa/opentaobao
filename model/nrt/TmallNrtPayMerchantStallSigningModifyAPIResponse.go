package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtPayMerchantStallSigningModifyAPIResponse 三级商户进件修改 API返回值
// tmall.nrt.pay.merchant.stall.signing.modify
//
// 三级商户进件修改
type TmallNrtPayMerchantStallSigningModifyAPIResponse struct {
	model.CommonResponse
	TmallNrtPayMerchantStallSigningModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtPayMerchantStallSigningModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtPayMerchantStallSigningModifyAPIResponseModel).Reset()
}

// TmallNrtPayMerchantStallSigningModifyAPIResponseModel is 三级商户进件修改 成功返回结果
type TmallNrtPayMerchantStallSigningModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_pay_merchant_stall_signing_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallNrtPayMerchantStallSigningModifyResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtPayMerchantStallSigningModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrtPayMerchantStallSigningModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtPayMerchantStallSigningModifyAPIResponse)
	},
}

// GetTmallNrtPayMerchantStallSigningModifyAPIResponse 从 sync.Pool 获取 TmallNrtPayMerchantStallSigningModifyAPIResponse
func GetTmallNrtPayMerchantStallSigningModifyAPIResponse() *TmallNrtPayMerchantStallSigningModifyAPIResponse {
	return poolTmallNrtPayMerchantStallSigningModifyAPIResponse.Get().(*TmallNrtPayMerchantStallSigningModifyAPIResponse)
}

// ReleaseTmallNrtPayMerchantStallSigningModifyAPIResponse 将 TmallNrtPayMerchantStallSigningModifyAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtPayMerchantStallSigningModifyAPIResponse(v *TmallNrtPayMerchantStallSigningModifyAPIResponse) {
	v.Reset()
	poolTmallNrtPayMerchantStallSigningModifyAPIResponse.Put(v)
}
