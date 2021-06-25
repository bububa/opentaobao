package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询商品 APIResponse
taobao.scitem.get

根据id查询商品
*/
type TaobaoScitemGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoScitemGetResponse `json:"taobao_scitem_get_response,omitempty"`
}

type TaobaoScitemGetResponse struct {

    // 后端商品
    ScItem   *ScItem `json:"sc_item,omitempty"`

}
