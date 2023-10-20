package qt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotssubscribegetAPIResponse 淘宝服务订购关系查询 API返回值
// taobao.ts.subscribe.get
//
// ts订购关系状态查询. 暂只支持1口价服务.
type TaobaotssubscribegetAPIResponse struct {
	model.CommonResponse
	TaobaotssubscribegetAPIResponseModel
}

// TaobaotssubscribegetAPIResponseModel is 淘宝服务订购关系查询 成功返回结果
type TaobaotssubscribegetAPIResponseModel struct {
	XMLName xml.Name `xml:"ts_subscribe_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订购关系对象
	ServiceSubscribe *ServiceSubscribe `json:"service_subscribe,omitempty" xml:"service_subscribe,omitempty"`
}
