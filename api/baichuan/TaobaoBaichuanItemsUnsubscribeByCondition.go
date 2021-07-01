package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

/* TaobaoBaichuanItemsUnsubscribeByCondition
根据条件删除订阅关系
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系 */
func TaobaoBaichuanItemsUnsubscribeByCondition(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest, session string) (*baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse, error) {
	var resp baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
