package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaominiappcloudfunctioninvoke 外部触发云函数
// taobao.miniapp.cloud.function.invoke
//
// 用户isv从外部触发聚石塔云函数的执行。
func Taobaominiappcloudfunctioninvoke(clt *core.SDKClient, req *miniapp.TaobaominiappcloudfunctioninvokeAPIRequest, session string) (*miniapp.TaobaominiappcloudfunctioninvokeAPIResponse, error) {
	var resp miniapp.TaobaominiappcloudfunctioninvokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
