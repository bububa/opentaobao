package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyShopUpdatestatusAPIResponse 更新渠道店营业状态 API返回值
// alibaba.aelophy.shop.updatestatus
//
// 更新渠道店营业状态
type AlibabaAelophyShopUpdatestatusAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyShopUpdatestatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyShopUpdatestatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyShopUpdatestatusAPIResponseModel).Reset()
}

// AlibabaAelophyShopUpdatestatusAPIResponseModel is 更新渠道店营业状态 成功返回结果
type AlibabaAelophyShopUpdatestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_shop_updatestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// api调用结果
	ApiResult *AlibabaAelophyShopUpdatestatusApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyShopUpdatestatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAelophyShopUpdatestatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyShopUpdatestatusAPIResponse)
	},
}

// GetAlibabaAelophyShopUpdatestatusAPIResponse 从 sync.Pool 获取 AlibabaAelophyShopUpdatestatusAPIResponse
func GetAlibabaAelophyShopUpdatestatusAPIResponse() *AlibabaAelophyShopUpdatestatusAPIResponse {
	return poolAlibabaAelophyShopUpdatestatusAPIResponse.Get().(*AlibabaAelophyShopUpdatestatusAPIResponse)
}

// ReleaseAlibabaAelophyShopUpdatestatusAPIResponse 将 AlibabaAelophyShopUpdatestatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyShopUpdatestatusAPIResponse(v *AlibabaAelophyShopUpdatestatusAPIResponse) {
	v.Reset()
	poolAlibabaAelophyShopUpdatestatusAPIResponse.Put(v)
}
