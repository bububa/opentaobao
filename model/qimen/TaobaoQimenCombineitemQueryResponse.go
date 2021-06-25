package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合货品关系查询接口 APIResponse
taobao.qimen.combineitem.query

组合货品关系查询
*/
type TaobaoQimenCombineitemQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenCombineitemQueryResponse `json:"taobao_qimen_combineitem_query_response,omitempty"`
}

type TaobaoQimenCombineitemQueryResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
