package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAdgroupHorizontalFindpage 查询单元分页列表
// taobao.universalbp.adgroup.horizontal.findpage
//
// 查询单元分页列表
func TaobaoUniversalbpAdgroupHorizontalFindpage(clt *core.SDKClient, req *simba.TaobaoUniversalbpAdgroupHorizontalFindpageAPIRequest, session string) (*simba.TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse, error) {
	var resp simba.TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
