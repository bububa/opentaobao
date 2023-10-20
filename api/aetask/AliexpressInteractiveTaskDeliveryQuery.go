package aetask

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetask"
)

// AliexpressInteractiveTaskDeliveryQuery AE互动任务投放
// aliexpress.interactive.task.delivery.query
//
// 将内部配置好的任务，如浏览商品，店铺投放给外部ISV
func AliexpressInteractiveTaskDeliveryQuery(clt *core.SDKClient, req *aetask.AliexpressInteractiveTaskDeliveryQueryAPIRequest, resp *aetask.AliexpressInteractiveTaskDeliveryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
