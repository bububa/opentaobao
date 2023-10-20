package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoqimentradeuseradd 添加奇门订单链路用户
// taobao.qimen.trade.user.add
//
// 添加奇门订单链路用户
func Taobaoqimentradeuseradd(clt *core.SDKClient, req *util.TaobaoqimentradeuseraddAPIRequest, session string) (*util.TaobaoqimentradeuseraddAPIResponse, error) {
	var resp util.TaobaoqimentradeuseraddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
