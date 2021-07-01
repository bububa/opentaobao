package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

/* AlibabaAlisportsPassportOauthAlipaygrant
阿里体育会员系统-支付宝授权接口
alibaba.alisports.passport.oauth.alipaygrant

开放给乐心运动使用的支付宝授权接口 */
func AlibabaAlisportsPassportOauthAlipaygrant(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportOauthAlipaygrantAPIRequest, session string) (*alisports.AlibabaAlisportsPassportOauthAlipaygrantAPIResponse, error) {
	var resp alisports.AlibabaAlisportsPassportOauthAlipaygrantAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
