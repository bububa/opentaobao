package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemmappingCreateAPIResponse 前后端商品映射接口 API返回值
// taobao.qimen.itemmapping.create
//
// 前后端商品映射
type TaobaoQimenItemmappingCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemmappingCreateAPIResponseModel
}

// TaobaoQimenItemmappingCreateAPIResponseModel is 前后端商品映射接口 成功返回结果
type TaobaoQimenItemmappingCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_itemmapping_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}
