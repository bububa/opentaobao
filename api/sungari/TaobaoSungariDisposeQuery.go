package sungari

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// Taobaosungaridisposequery 商品商家处置结果查询
// taobao.sungari.dispose.query
//
// 红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
func Taobaosungaridisposequery(clt *core.SDKClient, req *sungari.TaobaosungaridisposequeryAPIRequest, session string) (*sungari.TaobaosungaridisposequeryAPIResponse, error) {
	var resp sungari.TaobaosungaridisposequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
