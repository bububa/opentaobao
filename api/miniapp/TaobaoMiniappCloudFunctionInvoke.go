package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappCloudFunctionInvoke 外部触发云函数
// taobao.miniapp.cloud.function.invoke
//
// 用户isv从外部触发聚石塔云函数的执行。
func TaobaoMiniappCloudFunctionInvoke(clt *core.SDKClient, req *miniapp.TaobaoMiniappCloudFunctionInvokeAPIRequest, resp *miniapp.TaobaoMiniappCloudFunctionInvokeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
