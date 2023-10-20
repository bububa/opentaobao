package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemitempage 信息流查看商品列表
// taobao.feedflow.item.item.page
//
// 信息流查看商品列表
func Taobaofeedflowitemitempage(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemitempageAPIRequest, session string) (*feedflow.TaobaofeedflowitemitempageAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemitempageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
