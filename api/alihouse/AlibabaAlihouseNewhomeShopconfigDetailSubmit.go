package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigDetailSubmit 店铺配置信息接口
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
func AlibabaAlihouseNewhomeShopconfigDetailSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
