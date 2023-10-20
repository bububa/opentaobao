package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportaccountcheckmobile 阿里体育会员系统--手机号验证接口
// alibaba.alisports.passport.account.checkmobile
//
// 验证三方用户的手机号，获取对应的阿里体育会员id
func Alibabaalisportspassportaccountcheckmobile(clt *core.SDKClient, req *alisports.AlibabaalisportspassportaccountcheckmobileAPIRequest, session string) (*alisports.AlibabaalisportspassportaccountcheckmobileAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportaccountcheckmobileAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
