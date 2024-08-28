package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoCarrierdriverLocation 查询骑手当前位置
// alibaba.ele.fengniao.carrierdriver.location
//
// 查询骑手当前位置
func AlibabaEleFengniaoCarrierdriverLocation(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoCarrierdriverLocationAPIRequest, resp *logistic.AlibabaEleFengniaoCarrierdriverLocationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
