package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAuthAccountinfo 授权账号信息
// alibaba.alisports.passport.auth.accountinfo
//
// 获取体育用户OpenId\nick\avatar 三个信息
func AlibabaAlisportsPassportAuthAccountinfo(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAuthAccountinfoAPIRequest, resp *alisports.AlibabaAlisportsPassportAuthAccountinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
