package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanorderurlgetAPIResponse 百川订单详情 API返回值
// taobao.baichuan.orderurl.get
//
// 百川订单详情
type TaobaobaichuanorderurlgetAPIResponse struct {
	model.CommonResponse
	TaobaobaichuanorderurlgetAPIResponseModel
}

// TaobaobaichuanorderurlgetAPIResponseModel is 百川订单详情 成功返回结果
type TaobaobaichuanorderurlgetAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_orderurl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
