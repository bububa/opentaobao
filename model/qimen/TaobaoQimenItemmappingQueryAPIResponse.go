package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
前后端商品映射查询接口 API返回值 
taobao.qimen.itemmapping.query

前后端商品映射查询接口
*/
type TaobaoQimenItemmappingQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenItemmappingQueryAPIResponseModel
}

// 前后端商品映射查询接口 成功返回结果
type TaobaoQimenItemmappingQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_itemmapping_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *TaobaoQimenItemmappingQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}
