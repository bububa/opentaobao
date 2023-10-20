package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupRecommendProductAPIResponse 推品 API返回值
// alibaba.scbp.ad.group.recommend.product
//
// 推品
type AlibabaScbpAdGroupRecommendProductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupRecommendProductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupRecommendProductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupRecommendProductAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupRecommendProductAPIResponseModel is 推品 成功返回结果
type AlibabaScbpAdGroupRecommendProductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_recommend_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推品返回实体类
	ResultList []RecommendProductDto `json:"result_list,omitempty" xml:"result_list>recommend_product_dto,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupRecommendProductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolAlibabaScbpAdGroupRecommendProductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupRecommendProductAPIResponse)
	},
}

// GetAlibabaScbpAdGroupRecommendProductAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupRecommendProductAPIResponse
func GetAlibabaScbpAdGroupRecommendProductAPIResponse() *AlibabaScbpAdGroupRecommendProductAPIResponse {
	return poolAlibabaScbpAdGroupRecommendProductAPIResponse.Get().(*AlibabaScbpAdGroupRecommendProductAPIResponse)
}

// ReleaseAlibabaScbpAdGroupRecommendProductAPIResponse 将 AlibabaScbpAdGroupRecommendProductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupRecommendProductAPIResponse(v *AlibabaScbpAdGroupRecommendProductAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupRecommendProductAPIResponse.Put(v)
}
