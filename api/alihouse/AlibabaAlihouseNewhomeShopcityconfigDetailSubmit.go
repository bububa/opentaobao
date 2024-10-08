package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmit 城市配置信息接口
// alibaba.alihouse.newhome.shopcityconfig.detail.submit
//
// 上传城市配置信息
func AlibabaAlihouseNewhomeShopcityconfigDetailSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
