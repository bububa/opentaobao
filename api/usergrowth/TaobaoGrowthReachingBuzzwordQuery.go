package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// Taobaogrowthreachingbuzzwordquery 淘宝热词榜单数据查询接口
// taobao.growth.reaching.buzzword.query
//
// 查询淘宝热词榜单数据
func Taobaogrowthreachingbuzzwordquery(clt *core.SDKClient, req *usergrowth.TaobaogrowthreachingbuzzwordqueryAPIRequest, session string) (*usergrowth.TaobaogrowthreachingbuzzwordqueryAPIResponse, error) {
	var resp usergrowth.TaobaogrowthreachingbuzzwordqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
