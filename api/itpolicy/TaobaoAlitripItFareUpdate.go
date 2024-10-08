package itpolicy

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItFareUpdate 【国际机票自有政策】单条修改
// taobao.alitrip.it.fare.update
//
// 自有政策修改接口，可以根据fareId或outId修改，outId不唯一时，不能用outId修改。当外部政策id、出发城市、到达城市、出票航司任一有变化，或往返时是否允许混舱、文件编号、可混文件编号任一有变化，将删除老数据，产生一条新政策。
func TaobaoAlitripItFareUpdate(ctx context.Context, clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareUpdateAPIRequest, resp *itpolicy.TaobaoAlitripItFareUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
