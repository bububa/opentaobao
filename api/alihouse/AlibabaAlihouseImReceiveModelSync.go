package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseImReceiveModelSync IM承接方式同步
// alibaba.alihouse.im.receive.model.sync
//
// IM承接方式同步
func AlibabaAlihouseImReceiveModelSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseImReceiveModelSyncAPIRequest, resp *alihouse.AlibabaAlihouseImReceiveModelSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
