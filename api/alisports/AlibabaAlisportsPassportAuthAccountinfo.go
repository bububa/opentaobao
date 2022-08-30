package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAuthAccountinfo 授权账号信息
// alibaba.alisports.passport.auth.accountinfo
//
// 获取体育用户OpenId\nick\avatar 三个信息
func AlibabaAlisportsPassportAuthAccountinfo(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAuthAccountinfoAPIRequest, session string) (*alisports.AlibabaAlisportsPassportAuthAccountinfoAPIResponse, error) {
	var resp alisports.AlibabaAlisportsPassportAuthAccountinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
