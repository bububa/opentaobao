package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarFpcarGetcarNotify 门店通知用户提车
// tmall.car.fpcar.getcar.notify
//
// 提供给外部(大搜或其它合作方)的接口-门店通知用户提车
func TmallCarFpcarGetcarNotify(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallCarFpcarGetcarNotifyAPIRequest, resp *servicecenter.TmallCarFpcarGetcarNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
