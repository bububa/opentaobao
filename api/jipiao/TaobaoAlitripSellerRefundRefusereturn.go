package jipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jipiao"
)

// Taobaoalitripsellerrefundrefusereturn 【机票代理商】拒绝退票
// taobao.alitrip.seller.refund.refusereturn
//
// 拒绝退票
func Taobaoalitripsellerrefundrefusereturn(clt *core.SDKClient, req *jipiao.TaobaoalitripsellerrefundrefusereturnAPIRequest, session string) (*jipiao.TaobaoalitripsellerrefundrefusereturnAPIResponse, error) {
	var resp jipiao.TaobaoalitripsellerrefundrefusereturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
