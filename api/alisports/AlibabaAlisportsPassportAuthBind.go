package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportauthbind 授权绑定关系接口
// alibaba.alisports.passport.auth.bind
//
// 授权回调绑定关系接口，建立阿里体育openId和三方openId的绑定关系
func Alibabaalisportspassportauthbind(clt *core.SDKClient, req *alisports.AlibabaalisportspassportauthbindAPIRequest, session string) (*alisports.AlibabaalisportspassportauthbindAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportauthbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
