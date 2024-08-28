package itpolicy

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItFareAddrt 【国际机票自有政策】单条往返添加
// taobao.alitrip.it.fare.addrt
//
// 自有政策往返添加接口
func TaobaoAlitripItFareAddrt(ctx context.Context, clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareAddrtAPIRequest, resp *itpolicy.TaobaoAlitripItFareAddrtAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
