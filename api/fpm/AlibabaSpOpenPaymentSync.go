package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// Alibabaspopenpaymentsync 付款单同步
// alibaba.sp.open.payment.sync
//
// 新康众弹外同步付款数据
func Alibabaspopenpaymentsync(clt *core.SDKClient, req *fpm.AlibabaspopenpaymentsyncAPIRequest, session string) (*fpm.AlibabaspopenpaymentsyncAPIResponse, error) {
	var resp fpm.AlibabaspopenpaymentsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
