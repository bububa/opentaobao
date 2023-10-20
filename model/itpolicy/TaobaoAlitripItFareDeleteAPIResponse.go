package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareDeleteAPIResponse 【国际机票自有政策】单条删除 API返回值
// taobao.alitrip.it.fare.delete
//
// 自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败
type TaobaoAlitripItFareDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItFareDeleteAPIResponseModel
}

// TaobaoAlitripItFareDeleteAPIResponseModel is 【国际机票自有政策】单条删除 成功返回结果
type TaobaoAlitripItFareDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
}
