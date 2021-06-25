package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发布后端商品 APIResponse
taobao.scitem.add

发布后端商品
*/
type TaobaoScitemAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoScitemAddResponse `json:"taobao_scitem_add_response,omitempty"`
}

type TaobaoScitemAddResponse struct {

    // 后台商品信息
    ScItem   *ScItem `json:"sc_item,omitempty"`

}
