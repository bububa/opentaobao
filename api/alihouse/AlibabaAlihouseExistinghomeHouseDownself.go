package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseDownself 房源信息下架
// alibaba.alihouse.existinghome.house.downself
//
// 房源信息下架
func AlibabaAlihouseExistinghomeHouseDownself(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseDownselfAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseDownselfAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
