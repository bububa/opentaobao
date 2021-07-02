package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAccountTokenvalidate 阿里体育会员系统帐号登录注册token验证接口
// alibaba.alisports.passport.account.tokenvalidate
//
// 阿里体育会员系统帐号登录注册token验证接口
func AlibabaAlisportsPassportAccountTokenvalidate(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAccountTokenvalidateAPIRequest, session string) (*alisports.AlibabaAlisportsPassportAccountTokenvalidateAPIResponse, error) {
	var resp alisports.AlibabaAlisportsPassportAccountTokenvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
