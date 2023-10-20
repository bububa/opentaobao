package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaoqimentradeuserdelete 删除奇门订单链路用户
// taobao.qimen.trade.user.delete
//
// 删除奇门订单链路用户
func Taobaoqimentradeuserdelete(clt *core.SDKClient, req *util.TaobaoqimentradeuserdeleteAPIRequest, session string) (*util.TaobaoqimentradeuserdeleteAPIResponse, error) {
	var resp util.TaobaoqimentradeuserdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
