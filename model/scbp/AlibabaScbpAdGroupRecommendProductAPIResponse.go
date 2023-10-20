package scbp

import (
	"encoding/xml"

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

// AlibabaScbpAdGroupRecommendProductAPIResponseModel is 推品 成功返回结果
type AlibabaScbpAdGroupRecommendProductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_recommend_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推品返回实体类
	ResultList []RecommendProductDto `json:"result_list,omitempty" xml:"result_list>recommend_product_dto,omitempty"`
}
