package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合货品删除接口 APIResponse
taobao.qimen.combineitem.delete

组合货品删除
*/
type TaobaoQimenCombineitemDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenCombineitemDeleteResponse `json:"taobao_qimen_combineitem_delete_response,omitempty"`
}

type TaobaoQimenCombineitemDeleteResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
