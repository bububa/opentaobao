package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosGoodsBulkinputcspu 批量录入商品信息
// alibaba.mos.goods.bulkinputcspu
//
// 用于商品信息的批量导入到银泰商品中台
func AlibabaMosGoodsBulkinputcspu(clt *core.SDKClient, req *moscm.AlibabaMosGoodsBulkinputcspuAPIRequest, resp *moscm.AlibabaMosGoodsBulkinputcspuAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
