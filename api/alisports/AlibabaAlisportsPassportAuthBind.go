package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAuthBind 授权绑定关系接口
// alibaba.alisports.passport.auth.bind
//
// 授权回调绑定关系接口，建立阿里体育openId和三方openId的绑定关系
func AlibabaAlisportsPassportAuthBind(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAuthBindAPIRequest, resp *alisports.AlibabaAlisportsPassportAuthBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
