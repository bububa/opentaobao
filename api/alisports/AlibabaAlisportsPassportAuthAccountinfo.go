package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportauthaccountinfo 授权账号信息
// alibaba.alisports.passport.auth.accountinfo
//
// 获取体育用户OpenId\nick\avatar 三个信息
func Alibabaalisportspassportauthaccountinfo(clt *core.SDKClient, req *alisports.AlibabaalisportspassportauthaccountinfoAPIRequest, session string) (*alisports.AlibabaalisportspassportauthaccountinfoAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportauthaccountinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
