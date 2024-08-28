package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopconfigDetailSubmit 店铺配置信息接口
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
func AlibabaAlihouseNewhomeShopconfigDetailSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
