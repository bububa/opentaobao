package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabaguardaccessauth 鉴权
// alibaba.guard.access.auth
//
// 刷卡鉴权
func Alibabaguardaccessauth(clt *core.SDKClient, req *campus.AlibabaguardaccessauthAPIRequest, session string) (*campus.AlibabaguardaccessauthAPIResponse, error) {
	var resp campus.AlibabaguardaccessauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
