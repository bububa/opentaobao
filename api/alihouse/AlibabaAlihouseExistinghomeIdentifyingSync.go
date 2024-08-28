package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeIdentifyingSync 登陆标识信息同步
// alibaba.alihouse.existinghome.identifying.sync
//
// 登陆标识信息同步
func AlibabaAlihouseExistinghomeIdentifyingSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
