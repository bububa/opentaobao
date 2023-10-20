package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportaccountbindthirdid 阿里体育三方ID绑定接口
// alibaba.alisports.passport.account.bindthirdid
//
// 阿里体育三方ID绑定接口
func Alibabaalisportspassportaccountbindthirdid(clt *core.SDKClient, req *alisports.AlibabaalisportspassportaccountbindthirdidAPIRequest, session string) (*alisports.AlibabaalisportspassportaccountbindthirdidAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportaccountbindthirdidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
