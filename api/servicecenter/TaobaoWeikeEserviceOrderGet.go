package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoWeikeEserviceOrderGet 客服外包订单查询
// taobao.weike.eservice.order.get
//
// 用于客服外包中服务商查询订单列表
func TaobaoWeikeEserviceOrderGet(clt *core.SDKClient, req *servicecenter.TaobaoWeikeEserviceOrderGetAPIRequest, session string) (*servicecenter.TaobaoWeikeEserviceOrderGetAPIResponse, error) {
	var resp servicecenter.TaobaoWeikeEserviceOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
