package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流宝商品修改 APIResponse
taobao.wlb.item.update

修改物流宝商品信息
*/
type TaobaoWlbItemUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbItemUpdateResponse `json:"taobao_wlb_item_update_response,omitempty"`
}

type TaobaoWlbItemUpdateResponse struct {

    // 修改时间
    GmtModified   bool `json:"gmt_modified,omitempty"`

}
