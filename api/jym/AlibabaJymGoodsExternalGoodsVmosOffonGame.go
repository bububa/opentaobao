package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymGoodsExternalGoodsVmosOffonGame 基于游戏id临时上下架智能发布入口
// alibaba.jym.goods.external.goods.vmos.offon.game
//
// 基于游戏id临时上下架智能发布入口
func AlibabaJymGoodsExternalGoodsVmosOffonGame(clt *core.SDKClient, req *jym.AlibabaJymGoodsExternalGoodsVmosOffonGameAPIRequest, session string) (*jym.AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponse, error) {
	var resp jym.AlibabaJymGoodsExternalGoodsVmosOffonGameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
