package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateHotproductDownloadAPIResponse 联盟营销爆品下载接口 API返回值
// aliexpress.affiliate.hotproduct.download
//
// 查询联盟爆品API
type AliexpressAffiliateHotproductDownloadAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateHotproductDownloadAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateHotproductDownloadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateHotproductDownloadAPIResponseModel).Reset()
}

// AliexpressAffiliateHotproductDownloadAPIResponseModel is 联盟营销爆品下载接口 成功返回结果
type AliexpressAffiliateHotproductDownloadAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_hotproduct_download_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateHotproductDownloadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateHotproductDownloadAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateHotproductDownloadAPIResponse)
	},
}

// GetAliexpressAffiliateHotproductDownloadAPIResponse 从 sync.Pool 获取 AliexpressAffiliateHotproductDownloadAPIResponse
func GetAliexpressAffiliateHotproductDownloadAPIResponse() *AliexpressAffiliateHotproductDownloadAPIResponse {
	return poolAliexpressAffiliateHotproductDownloadAPIResponse.Get().(*AliexpressAffiliateHotproductDownloadAPIResponse)
}

// ReleaseAliexpressAffiliateHotproductDownloadAPIResponse 将 AliexpressAffiliateHotproductDownloadAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateHotproductDownloadAPIResponse(v *AliexpressAffiliateHotproductDownloadAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateHotproductDownloadAPIResponse.Put(v)
}
