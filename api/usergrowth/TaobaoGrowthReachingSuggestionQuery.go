package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaogrowthreachingsuggestionquery 厂商生态推荐信息查询
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
func Taobaogrowthreachingsuggestionquery(clt *core.SDKClient, req *usergrowth.TaobaogrowthreachingsuggestionqueryAPIRequest, session string) (*usergrowth.TaobaogrowthreachingsuggestionqueryAPIResponse, error) {
	var resp usergrowth.TaobaogrowthreachingsuggestionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
