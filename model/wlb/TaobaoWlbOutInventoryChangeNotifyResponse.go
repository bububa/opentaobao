package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部库存变化通知（企业物流用户使用） APIResponse
taobao.wlb.out.inventory.change.notify

拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
*/
type TaobaoWlbOutInventoryChangeNotifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbOutInventoryChangeNotifyResponse `json:"taobao_wlb_out_inventory_change_notify_response,omitempty"`
}

type TaobaoWlbOutInventoryChangeNotifyResponse struct {

    // 库存变化通知成功时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
