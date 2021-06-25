package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品查询接口 APIResponse
taobao.qimen.singleitem.query

商品查询接口
*/
type TaobaoQimenSingleitemQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenSingleitemQueryResponse `json:"taobao_qimen_singleitem_query_response,omitempty"`
}

type TaobaoQimenSingleitemQueryResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
