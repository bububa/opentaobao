package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIsvPageQuery isv离线会员数据分页查询
// alibaba.member.isv.page.query
//
// isv离线会员数据分页查询
func AlibabaMemberIsvPageQuery(clt *core.SDKClient, req *alimember.AlibabaMemberIsvPageQueryAPIRequest, resp *alimember.AlibabaMemberIsvPageQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
