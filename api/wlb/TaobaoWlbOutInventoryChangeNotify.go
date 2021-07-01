package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

/* TaobaoWlbOutInventoryChangeNotify
外部库存变化通知（企业物流用户使用）
taobao.wlb.out.inventory.change.notify

拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。 */
func TaobaoWlbOutInventoryChangeNotify(clt *core.SDKClient, req *wlb.TaobaoWlbOutInventoryChangeNotifyAPIRequest, session string) (*wlb.TaobaoWlbOutInventoryChangeNotifyAPIResponse, error) {
	var resp wlb.TaobaoWlbOutInventoryChangeNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
