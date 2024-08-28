package itpolicy

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItFareAddow 【国际机票自有政策】单条单程添加
// taobao.alitrip.it.fare.addow
//
// 自有政策单程添加接口，重复的老数据会被删除，重复判断规则同excel
func TaobaoAlitripItFareAddow(ctx context.Context, clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareAddowAPIRequest, resp *itpolicy.TaobaoAlitripItFareAddowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
