package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsRecommendFeedGetAPIResponse 获取推荐商品信息流接口 API返回值
// aliexpress.ds.recommend.feed.get
//
// 获取推荐商品信息流
type AliexpressDsRecommendFeedGetAPIResponse struct {
	model.CommonResponse
	AliexpressDsRecommendFeedGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressDsRecommendFeedGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressDsRecommendFeedGetAPIResponseModel).Reset()
}

// AliexpressDsRecommendFeedGetAPIResponseModel is 获取推荐商品信息流接口 成功返回结果
type AliexpressDsRecommendFeedGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_recommend_feed_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// System Error
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// result object
	Result *TrafficProductResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressDsRecommendFeedGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RspMsg = ""
	m.RspCode = ""
	m.Result = nil
}

var poolAliexpressDsRecommendFeedGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressDsRecommendFeedGetAPIResponse)
	},
}

// GetAliexpressDsRecommendFeedGetAPIResponse 从 sync.Pool 获取 AliexpressDsRecommendFeedGetAPIResponse
func GetAliexpressDsRecommendFeedGetAPIResponse() *AliexpressDsRecommendFeedGetAPIResponse {
	return poolAliexpressDsRecommendFeedGetAPIResponse.Get().(*AliexpressDsRecommendFeedGetAPIResponse)
}

// ReleaseAliexpressDsRecommendFeedGetAPIResponse 将 AliexpressDsRecommendFeedGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressDsRecommendFeedGetAPIResponse(v *AliexpressDsRecommendFeedGetAPIResponse) {
	v.Reset()
	poolAliexpressDsRecommendFeedGetAPIResponse.Put(v)
}
