package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeLineSync 环线数据同步
// alibaba.alihouse.newhome.line.sync
//
// 环线数据同步
func AlibabaAlihouseNewhomeLineSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLineSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeLineSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
