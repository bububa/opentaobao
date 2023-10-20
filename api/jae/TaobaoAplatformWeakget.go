package jae

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jae"
)

// Taobaoaplatformweakget 活动平台弱登录接口
// taobao.aplatform.weakget
//
// 无线活动平台的开放接口，提供商品信息等的读操作
func Taobaoaplatformweakget(clt *core.SDKClient, req *jae.TaobaoaplatformweakgetAPIRequest, session string) (*jae.TaobaoaplatformweakgetAPIResponse, error) {
	var resp jae.TaobaoaplatformweakgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
