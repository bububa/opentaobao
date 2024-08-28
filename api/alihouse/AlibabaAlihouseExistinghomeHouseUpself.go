package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseUpself 房源上架
// alibaba.alihouse.existinghome.house.upself
//
// 房源信息上架
func AlibabaAlihouseExistinghomeHouseUpself(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseUpselfAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseUpselfAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
