package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAdzoneHorizontalFindpage 查看资源包列表
// taobao.universalbp.adzone.horizontal.findpage
//
// 查看已存在的计划上设置的资源包列表
func TaobaoUniversalbpAdzoneHorizontalFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest, resp *simba.TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
