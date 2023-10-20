package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmallitemcentersubscribequeryAPIResponse 天猫服务订购信息查询接口 API返回值
// tmall.mallitemcenter.subscribe.query
//
// 查询商家服务订购信息
type TmallmallitemcentersubscribequeryAPIResponse struct {
	model.CommonResponse
	TmallmallitemcentersubscribequeryAPIResponseModel
}

// TmallmallitemcentersubscribequeryAPIResponseModel is 天猫服务订购信息查询接口 成功返回结果
type TmallmallitemcentersubscribequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mallitemcenter_subscribe_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallmallitemcentersubscribequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
