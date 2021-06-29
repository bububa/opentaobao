package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射接口 API返回值 
taobao.qimen.itemmapping.create

前后端商品映射
*/
type TaobaoQimenItemmappingCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQimenItemmappingCreateResponse
}

// 前后端商品映射接口 成功返回结果
type TaobaoQimenItemmappingCreateResponse struct {
    XMLName xml.Name `xml:"qimen_itemmapping_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *ResponseDO `json:"response,omitempty" xml:"response,omitempty"`
}
