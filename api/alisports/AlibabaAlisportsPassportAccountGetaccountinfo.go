package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportaccountgetaccountinfo 获取会员信息
// alibaba.alisports.passport.account.getaccountinfo
//
// 获取阿里体育会员信息
func Alibabaalisportspassportaccountgetaccountinfo(clt *core.SDKClient, req *alisports.AlibabaalisportspassportaccountgetaccountinfoAPIRequest, session string) (*alisports.AlibabaalisportspassportaccountgetaccountinfoAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportaccountgetaccountinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
