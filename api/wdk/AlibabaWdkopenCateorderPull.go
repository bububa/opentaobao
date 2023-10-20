package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkopenCateorderPull 商户回传餐饮加工单状态
// alibaba.wdkopen.cateorder.pull
//
// 商户回传餐饮加工单状态
func AlibabaWdkopenCateorderPull(clt *core.SDKClient, req *wdk.AlibabaWdkopenCateorderPullAPIRequest, resp *wdk.AlibabaWdkopenCateorderPullAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
