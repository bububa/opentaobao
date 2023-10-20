package sungari

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// TaobaoSungariDisposeQuery 商品商家处置结果查询
// taobao.sungari.dispose.query
//
// 红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
func TaobaoSungariDisposeQuery(clt *core.SDKClient, req *sungari.TaobaoSungariDisposeQueryAPIRequest, resp *sungari.TaobaoSungariDisposeQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
