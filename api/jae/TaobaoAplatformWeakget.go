package jae

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jae"
)

// TaobaoAplatformWeakget 活动平台弱登录接口
// taobao.aplatform.weakget
//
// 无线活动平台的开放接口，提供商品信息等的读操作
func TaobaoAplatformWeakget(clt *core.SDKClient, req *jae.TaobaoAplatformWeakgetAPIRequest, resp *jae.TaobaoAplatformWeakgetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
