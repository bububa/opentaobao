package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountGetaccountinfo 获取会员信息
// alibaba.alisports.passport.account.getaccountinfo
//
// 获取阿里体育会员信息
func AlibabaAlisportsPassportAccountGetaccountinfo(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountGetaccountinfoAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountGetaccountinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
