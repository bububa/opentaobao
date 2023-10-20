package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountBindthirdid 阿里体育三方ID绑定接口
// alibaba.alisports.passport.account.bindthirdid
//
// 阿里体育三方ID绑定接口
func AlibabaAlisportsPassportAccountBindthirdid(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountBindthirdidAPIRequest, session string) (*alisports.AlibabaAlisportsPassportAccountBindthirdidAPIResponse, error) {
	var resp alisports.AlibabaAlisportsPassportAccountBindthirdidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
