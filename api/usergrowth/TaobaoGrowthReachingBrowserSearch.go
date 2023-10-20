package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaogrowthreachingbrowsersearch 查询搜索关联
// taobao.growth.reaching.browser.search
//
// 查询搜索关联
func Taobaogrowthreachingbrowsersearch(clt *core.SDKClient, req *usergrowth.TaobaogrowthreachingbrowsersearchAPIRequest, session string) (*usergrowth.TaobaogrowthreachingbrowsersearchAPIResponse, error) {
	var resp usergrowth.TaobaogrowthreachingbrowsersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
