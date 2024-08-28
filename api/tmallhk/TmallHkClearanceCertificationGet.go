package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceCertificationGet 获取订单清关材料实名信息
// tmall.hk.clearance.certification.get
//
// 获取订单清关材料实名信息
func TmallHkClearanceCertificationGet(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallHkClearanceCertificationGetAPIRequest, resp *tmallhk.TmallHkClearanceCertificationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
