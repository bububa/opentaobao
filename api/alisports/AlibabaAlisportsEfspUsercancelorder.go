package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsEfspUsercancelorder 用户取消订单
// alibaba.alisports.efsp.usercancelorder
//
// 用户取消订单
func AlibabaAlisportsEfspUsercancelorder(clt *core.SDKClient, req *alisports.AlibabaAlisportsEfspUsercancelorderAPIRequest, session string) (*alisports.AlibabaAlisportsEfspUsercancelorderAPIResponse, error) {
	var resp alisports.AlibabaAlisportsEfspUsercancelorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
