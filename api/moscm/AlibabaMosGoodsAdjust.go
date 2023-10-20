package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosGoodsAdjust 调整库存
// alibaba.mos.goods.adjust
//
// 库存调整接口
func AlibabaMosGoodsAdjust(clt *core.SDKClient, req *moscm.AlibabaMosGoodsAdjustAPIRequest, resp *moscm.AlibabaMosGoodsAdjustAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
