package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrokerPointsSync 经纪人积分同步
// alibaba.alihouse.existinghome.broker.points.sync
//
// 经纪人积分
func AlibabaAlihouseExistinghomeBrokerPointsSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrokerPointsSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrokerPointsSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
