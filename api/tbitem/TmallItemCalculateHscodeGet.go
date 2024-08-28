package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemCalculateHscodeGet 算法获取hscode
// tmall.item.calculate.hscode.get
//
// 算法获取hscode
func TmallItemCalculateHscodeGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemCalculateHscodeGetAPIRequest, resp *tbitem.TmallItemCalculateHscodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
