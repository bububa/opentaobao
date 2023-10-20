package mei

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallmeicrmcallbackpointchangeAPIResponse 品牌积分变更回调API API返回值
// tmall.mei.crm.callback.point.change
//
// 线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
type TmallmeicrmcallbackpointchangeAPIResponse struct {
	model.CommonResponse
	TmallmeicrmcallbackpointchangeAPIResponseModel
}

// TmallmeicrmcallbackpointchangeAPIResponseModel is 品牌积分变更回调API 成功返回结果
type TmallmeicrmcallbackpointchangeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mei_crm_callback_point_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
