package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaoweikeeservicesubusersget 客服外包订单分配的商家子账号列表
// taobao.weike.eservice.subusers.get
//
// 获取客服外包订单分配的商家子账号列表，以及授权状态
func Taobaoweikeeservicesubusersget(clt *core.SDKClient, req *servicecenter.TaobaoweikeeservicesubusersgetAPIRequest, session string) (*servicecenter.TaobaoweikeeservicesubusersgetAPIResponse, error) {
	var resp servicecenter.TaobaoweikeeservicesubusersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
