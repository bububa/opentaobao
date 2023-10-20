package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAdgroupHorizontalFindpage 查询单元分页列表
// taobao.universalbp.adgroup.horizontal.findpage
//
// 查询单元分页列表
func TaobaoUniversalbpAdgroupHorizontalFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest, resp *simba.TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
