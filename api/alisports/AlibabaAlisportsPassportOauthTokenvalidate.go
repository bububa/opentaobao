package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportoauthtokenvalidate 阿里体育会员系统--获取登录信息接口
// alibaba.alisports.passport.oauth.tokenvalidate
//
// 阿里体育会员系统--获取登录信息接口
func Alibabaalisportspassportoauthtokenvalidate(clt *core.SDKClient, req *alisports.AlibabaalisportspassportoauthtokenvalidateAPIRequest, session string) (*alisports.AlibabaalisportspassportoauthtokenvalidateAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportoauthtokenvalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
