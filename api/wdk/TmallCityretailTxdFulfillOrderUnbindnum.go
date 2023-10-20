package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Tmallcityretailtxdfulfillorderunbindnum 淘鲜达虚拟号服务解绑接口
// tmall.cityretail.txd.fulfill.order.unbindnum
//
// 淘鲜达虚拟号解绑服务接口，通过订阅关系id进行解绑。
func Tmallcityretailtxdfulfillorderunbindnum(clt *core.SDKClient, req *wdk.TmallcityretailtxdfulfillorderunbindnumAPIRequest, session string) (*wdk.TmallcityretailtxdfulfillorderunbindnumAPIResponse, error) {
	var resp wdk.TmallcityretailtxdfulfillorderunbindnumAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
