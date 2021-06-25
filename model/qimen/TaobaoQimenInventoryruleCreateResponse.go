package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
渠道间库存规则设置接口 APIResponse
taobao.qimen.inventoryrule.create

渠道间库存规则设置
*/
type TaobaoQimenInventoryruleCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenInventoryruleCreateResponse `json:"taobao_qimen_inventoryrule_create_response,omitempty"`
}

type TaobaoQimenInventoryruleCreateResponse struct {

    // 
    Response   *ResponseDO `json:"response,omitempty"`

}
