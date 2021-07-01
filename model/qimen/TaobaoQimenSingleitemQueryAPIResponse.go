package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品查询接口 API返回值 
taobao.qimen.singleitem.query

商品查询接口
*/
type TaobaoQimenSingleitemQueryAPIResponse struct {
    model.CommonResponse
    TaobaoQimenSingleitemQueryAPIResponseModel
}

// 商品查询接口 成功返回结果
type TaobaoQimenSingleitemQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_singleitem_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}
