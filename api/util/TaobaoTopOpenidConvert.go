package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopOpenidConvert 混淆nick转openid
// taobao.top.openid.convert
//
// 混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
func TaobaoTopOpenidConvert(clt *core.SDKClient, req *util.TaobaoTopOpenidConvertAPIRequest, resp *util.TaobaoTopOpenidConvertAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
