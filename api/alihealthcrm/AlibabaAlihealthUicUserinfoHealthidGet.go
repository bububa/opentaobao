package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthuicuserinfohealthidget 获取健康id
// alibaba.alihealth.uic.userinfo.healthid.get
//
// 根据支付宝用户ID获取用户健康ID
func Alibabaalihealthuicuserinfohealthidget(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthuicuserinfohealthidgetAPIRequest, session string) (*alihealthcrm.AlibabaalihealthuicuserinfohealthidgetAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthuicuserinfohealthidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
