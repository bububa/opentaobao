package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanItemsUnsubscribeByCondition 根据条件删除订阅关系
// taobao.baichuan.items.unsubscribe.by.condition
//
// 根据条件删除订阅关系
func TaobaoBaichuanItemsUnsubscribeByCondition(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest, resp *baichuan.TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
