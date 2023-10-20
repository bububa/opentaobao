package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvSpuSearch spu搜索接口
// alibaba.idle.isv.spu.search
//
// 搜索的品牌和型号，供服务商进行选择
func AlibabaIdleIsvSpuSearch(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvSpuSearchAPIRequest, resp *idleisv.AlibabaIdleIsvSpuSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
