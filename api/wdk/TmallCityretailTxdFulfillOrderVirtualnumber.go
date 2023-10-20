package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Tmallcityretailtxdfulfillordervirtualnumber 淘鲜达虚拟号服务接口
// tmall.cityretail.txd.fulfill.order.virtualnumber
//
// 虚拟小号绑定接口，只有开通了虚拟号服务的淘鲜达商家能绑定。
func Tmallcityretailtxdfulfillordervirtualnumber(clt *core.SDKClient, req *wdk.TmallcityretailtxdfulfillordervirtualnumberAPIRequest, session string) (*wdk.TmallcityretailtxdfulfillordervirtualnumberAPIResponse, error) {
	var resp wdk.TmallcityretailtxdfulfillordervirtualnumberAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
