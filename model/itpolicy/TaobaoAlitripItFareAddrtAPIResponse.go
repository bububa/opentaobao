package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripitfareaddrtAPIResponse 【国际机票自有政策】单条往返添加 API返回值
// taobao.alitrip.it.fare.addrt
//
// 自有政策往返添加接口
type TaobaoalitripitfareaddrtAPIResponse struct {
	model.CommonResponse
	TaobaoalitripitfareaddrtAPIResponseModel
}

// TaobaoalitripitfareaddrtAPIResponseModel is 【国际机票自有政策】单条往返添加 成功返回结果
type TaobaoalitripitfareaddrtAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_addrt_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 运价id
	FareId int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`
}
