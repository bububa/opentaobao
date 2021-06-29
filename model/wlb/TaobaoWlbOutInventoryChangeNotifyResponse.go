package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外部库存变化通知（企业物流用户使用） APIResponse
taobao.wlb.out.inventory.change.notify

拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
*/
type TaobaoWlbOutInventoryChangeNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOutInventoryChangeNotifyResponse
}

type TaobaoWlbOutInventoryChangeNotifyResponse struct {
    XMLName xml.Name `xml:"wlb_out_inventory_change_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 库存变化通知成功时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`

    
}
