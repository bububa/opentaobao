package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtpaymerchantfundstypemodify 修改摊位分账类型
// tmall.nrt.pay.merchant.fundstype.modify
//
// 修改摊位分账类型
func Tmallnrtpaymerchantfundstypemodify(clt *core.SDKClient, req *nrt.TmallnrtpaymerchantfundstypemodifyAPIRequest, session string) (*nrt.TmallnrtpaymerchantfundstypemodifyAPIResponse, error) {
	var resp nrt.TmallnrtpaymerchantfundstypemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
