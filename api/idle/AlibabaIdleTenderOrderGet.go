package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderorderget 暗拍读取订单
// alibaba.idle.tender.order.get
//
// 查询省心卖暗拍项目订单
func Alibabaidletenderorderget(clt *core.SDKClient, req *idle.AlibabaidletenderordergetAPIRequest, session string) (*idle.AlibabaidletenderordergetAPIResponse, error) {
	var resp idle.AlibabaidletenderordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
