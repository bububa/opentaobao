package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateFeaturedpromoGetAPIResponse 联盟主题推广活动信息获取 API返回值
// aliexpress.affiliate.featuredpromo.get
//
// 获取联盟主题推广活动信息
type AliexpressAffiliateFeaturedpromoGetAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateFeaturedpromoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateFeaturedpromoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateFeaturedpromoGetAPIResponseModel).Reset()
}

// AliexpressAffiliateFeaturedpromoGetAPIResponseModel is 联盟主题推广活动信息获取 成功返回结果
type AliexpressAffiliateFeaturedpromoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_featuredpromo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *TrafficFeaturedPromoResponseDto `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateFeaturedpromoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateFeaturedpromoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateFeaturedpromoGetAPIResponse)
	},
}

// GetAliexpressAffiliateFeaturedpromoGetAPIResponse 从 sync.Pool 获取 AliexpressAffiliateFeaturedpromoGetAPIResponse
func GetAliexpressAffiliateFeaturedpromoGetAPIResponse() *AliexpressAffiliateFeaturedpromoGetAPIResponse {
	return poolAliexpressAffiliateFeaturedpromoGetAPIResponse.Get().(*AliexpressAffiliateFeaturedpromoGetAPIResponse)
}

// ReleaseAliexpressAffiliateFeaturedpromoGetAPIResponse 将 AliexpressAffiliateFeaturedpromoGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateFeaturedpromoGetAPIResponse(v *AliexpressAffiliateFeaturedpromoGetAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateFeaturedpromoGetAPIResponse.Put(v)
}
