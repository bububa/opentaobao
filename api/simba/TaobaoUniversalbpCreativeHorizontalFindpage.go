package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCreativeHorizontalFindpage 横向管理创意分页查询
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
func TaobaoUniversalbpCreativeHorizontalFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest, resp *simba.TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
