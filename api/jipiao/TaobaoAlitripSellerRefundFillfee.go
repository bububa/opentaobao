package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripsellerrefundfillfee 机票代理商】回填手续费
// taobao.alitrip.seller.refund.fillfee
//
// 回填手续费
func Taobaoalitripsellerrefundfillfee(clt *core.SDKClient, req *jipiao.TaobaoalitripsellerrefundfillfeeAPIRequest, session string) (*jipiao.TaobaoalitripsellerrefundfillfeeAPIResponse, error) {
	var resp jipiao.TaobaoalitripsellerrefundfillfeeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
