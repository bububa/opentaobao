package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeDynamicdataSubmit 提交小区动态信息
// alibaba.alihouse.newhome.dynamicdata.submit
//
// 提交小区动态信息
func AlibabaAlihouseNewhomeDynamicdataSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
