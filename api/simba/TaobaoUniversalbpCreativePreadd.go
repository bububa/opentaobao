package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCreativePreadd 创建单品创意前置信息
// taobao.universalbp.creative.preadd
//
// 用于关键词场景创建单品创意前使用
func TaobaoUniversalbpCreativePreadd(clt *core.SDKClient, req *simba.TaobaoUniversalbpCreativePreaddAPIRequest, resp *simba.TaobaoUniversalbpCreativePreaddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
