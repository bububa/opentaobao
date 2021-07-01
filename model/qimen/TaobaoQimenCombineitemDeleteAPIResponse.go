package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合货品删除接口 API返回值 
taobao.qimen.combineitem.delete

组合货品删除
*/
type TaobaoQimenCombineitemDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoQimenCombineitemDeleteAPIResponseModel
}

// 组合货品删除接口 成功返回结果
type TaobaoQimenCombineitemDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_combineitem_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}
