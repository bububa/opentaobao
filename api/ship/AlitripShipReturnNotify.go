package ship

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ship"
)

// AlitripShipReturnNotify 船票退票退款回填接口
// alitrip.ship.return.notify
//
// 此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。
func AlitripShipReturnNotify(ctx context.Context, clt *core.SDKClient, req *ship.AlitripShipReturnNotifyAPIRequest, resp *ship.AlitripShipReturnNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
