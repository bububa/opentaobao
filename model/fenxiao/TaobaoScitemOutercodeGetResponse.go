package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据outerCode查询商品 APIResponse
taobao.scitem.outercode.get

根据outerCode查询商品
*/
type TaobaoScitemOutercodeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoScitemOutercodeGetResponse `json:"taobao_scitem_outercode_get_response,omitempty"`
}

type TaobaoScitemOutercodeGetResponse struct {

    // 后台商品
    ScItem   *ScItem `json:"sc_item,omitempty"`

}
