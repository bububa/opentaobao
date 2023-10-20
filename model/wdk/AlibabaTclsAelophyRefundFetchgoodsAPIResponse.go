package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundFetchgoodsAPIResponse saas 售后逆向 商户发起逆向取货 API返回值
// alibaba.tcls.aelophy.refund.fetchgoods
//
// saas 售后逆向 商户发起逆向取货
type AlibabaTclsAelophyRefundFetchgoodsAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyRefundFetchgoodsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundFetchgoodsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAelophyRefundFetchgoodsAPIResponseModel).Reset()
}

// AlibabaTclsAelophyRefundFetchgoodsAPIResponseModel is saas 售后逆向 商户发起逆向取货 成功返回结果
type AlibabaTclsAelophyRefundFetchgoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_fetchgoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabaTclsAelophyRefundFetchgoodsApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAelophyRefundFetchgoodsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTclsAelophyRefundFetchgoodsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAelophyRefundFetchgoodsAPIResponse)
	},
}

// GetAlibabaTclsAelophyRefundFetchgoodsAPIResponse 从 sync.Pool 获取 AlibabaTclsAelophyRefundFetchgoodsAPIResponse
func GetAlibabaTclsAelophyRefundFetchgoodsAPIResponse() *AlibabaTclsAelophyRefundFetchgoodsAPIResponse {
	return poolAlibabaTclsAelophyRefundFetchgoodsAPIResponse.Get().(*AlibabaTclsAelophyRefundFetchgoodsAPIResponse)
}

// ReleaseAlibabaTclsAelophyRefundFetchgoodsAPIResponse 将 AlibabaTclsAelophyRefundFetchgoodsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAelophyRefundFetchgoodsAPIResponse(v *AlibabaTclsAelophyRefundFetchgoodsAPIResponse) {
	v.Reset()
	poolAlibabaTclsAelophyRefundFetchgoodsAPIResponse.Put(v)
}
