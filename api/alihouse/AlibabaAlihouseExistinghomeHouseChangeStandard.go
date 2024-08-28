package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseChangeStandard 委托房源变更标准房源
// alibaba.alihouse.existinghome.house.change.standard
//
// 委托房源变更标准房源
func AlibabaAlihouseExistinghomeHouseChangeStandard(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseChangeStandardAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseChangeStandardAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
