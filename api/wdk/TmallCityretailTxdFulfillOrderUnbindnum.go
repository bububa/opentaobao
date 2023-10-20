package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TmallCityretailTxdFulfillOrderUnbindnum 淘鲜达虚拟号服务解绑接口
// tmall.cityretail.txd.fulfill.order.unbindnum
//
// 淘鲜达虚拟号解绑服务接口，通过订阅关系id进行解绑。
func TmallCityretailTxdFulfillOrderUnbindnum(clt *core.SDKClient, req *wdk.TmallCityretailTxdFulfillOrderUnbindnumAPIRequest, session string) (*wdk.TmallCityretailTxdFulfillOrderUnbindnumAPIResponse, error) {
	var resp wdk.TmallCityretailTxdFulfillOrderUnbindnumAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
