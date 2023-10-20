package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpadgrouphorizontalfindpage 查询单元分页列表
// taobao.universalbp.adgroup.horizontal.findpage
//
// 查询单元分页列表
func Taobaouniversalbpadgrouphorizontalfindpage(clt *core.SDKClient, req *simba.TaobaouniversalbpadgrouphorizontalfindpageAPIRequest, session string) (*simba.TaobaouniversalbpadgrouphorizontalfindpageAPIResponse, error) {
	var resp simba.TaobaouniversalbpadgrouphorizontalfindpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
