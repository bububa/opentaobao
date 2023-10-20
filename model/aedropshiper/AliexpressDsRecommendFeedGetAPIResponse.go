package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdsrecommendfeedgetAPIResponse 获取推荐商品信息流接口 API返回值
// aliexpress.ds.recommend.feed.get
//
// 获取推荐商品信息流
type AliexpressdsrecommendfeedgetAPIResponse struct {
	model.CommonResponse
	AliexpressdsrecommendfeedgetAPIResponseModel
}

// AliexpressdsrecommendfeedgetAPIResponseModel is 获取推荐商品信息流接口 成功返回结果
type AliexpressdsrecommendfeedgetAPIResponseModel struct {
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
