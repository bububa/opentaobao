package aetask

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aetask"
)

// Aliexpressinteractivetaskdeliveryquery AE互动任务投放
// aliexpress.interactive.task.delivery.query
//
// 将内部配置好的任务，如浏览商品，店铺投放给外部ISV
func Aliexpressinteractivetaskdeliveryquery(clt *core.SDKClient, req *aetask.AliexpressinteractivetaskdeliveryqueryAPIRequest, session string) (*aetask.AliexpressinteractivetaskdeliveryqueryAPIResponse, error) {
	var resp aetask.AliexpressinteractivetaskdeliveryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
