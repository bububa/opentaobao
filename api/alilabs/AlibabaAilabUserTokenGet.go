package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabusertokenget 三方账号获取 token
// alibaba.ailab.user.token.get
//
// inside 设备的三方 app，通过 extId、schema 生成 token
func Alibabaailabusertokenget(clt *core.SDKClient, req *alilabs.AlibabaailabusertokengetAPIRequest, session string) (*alilabs.AlibabaailabusertokengetAPIResponse, error) {
	var resp alilabs.AlibabaailabusertokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
