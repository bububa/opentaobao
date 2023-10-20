package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// Alibabaalisportspassportauthunbind 三方关系解绑接口
// alibaba.alisports.passport.auth.unbind
//
// 解除阿里体育openId和三方合作方的openId的绑定关系
func Alibabaalisportspassportauthunbind(clt *core.SDKClient, req *alisports.AlibabaalisportspassportauthunbindAPIRequest, session string) (*alisports.AlibabaalisportspassportauthunbindAPIResponse, error) {
	var resp alisports.AlibabaalisportspassportauthunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
