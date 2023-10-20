package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolcreateactivityAPIResponse 添加商品池活动 API返回值
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
type AlibabahmmarketingitempoolcreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitempoolcreateactivityAPIResponseModel
}

// AlibabahmmarketingitempoolcreateactivityAPIResponseModel is 添加商品池活动 成功返回结果
type AlibabahmmarketingitempoolcreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
