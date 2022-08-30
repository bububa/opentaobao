package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

// AlibabaAlisportsPassportAuthBind 授权绑定关系接口
// alibaba.alisports.passport.auth.bind
//
// 授权回调绑定关系接口，建立阿里体育openId和三方openId的绑定关系
func AlibabaAlisportsPassportAuthBind(clt *core.SDKClient, req *alisports.AlibabaAlisportsPassportAuthBindAPIRequest, session string) (*alisports.AlibabaAlisportsPassportAuthBindAPIResponse, error) {
	var resp alisports.AlibabaAlisportsPassportAuthBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
