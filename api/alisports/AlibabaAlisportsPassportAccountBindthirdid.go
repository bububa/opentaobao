package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountBindthirdid 阿里体育三方ID绑定接口
// alibaba.alisports.passport.account.bindthirdid
//
// 阿里体育三方ID绑定接口
func AlibabaAlisportsPassportAccountBindthirdid(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountBindthirdidAPIRequest, resp *alisports.AlibabaAlisportsPassportAccountBindthirdidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
