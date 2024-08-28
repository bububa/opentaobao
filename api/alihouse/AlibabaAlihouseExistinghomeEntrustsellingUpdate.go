package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeEntrustsellingUpdate 管家状态及房源信息接口
// alibaba.alihouse.existinghome.entrustselling.update
//
// 管家状态及房源信息接口
func AlibabaAlihouseExistinghomeEntrustsellingUpdate(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeEntrustsellingUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
