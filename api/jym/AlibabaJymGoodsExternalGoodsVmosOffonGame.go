package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymGoodsExternalGoodsVmosOffonGame 基于游戏id临时上下架智能发布入口
// alibaba.jym.goods.external.goods.vmos.offon.game
//
// 基于游戏id临时上下架智能发布入口
func AlibabaJymGoodsExternalGoodsVmosOffonGame(clt *core.SDKClient, req *jym.AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest, resp *jym.AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
