package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeBusinessSync 商圈数据同步
// alibaba.alihouse.newhome.business.sync
//
// 商圈数据同步
func AlibabaAlihouseNewhomeBusinessSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeBusinessSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeBusinessSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
