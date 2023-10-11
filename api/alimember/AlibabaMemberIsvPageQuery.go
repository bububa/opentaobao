package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIsvPageQuery isv离线会员数据分页查询
// alibaba.member.isv.page.query
//
// isv离线会员数据分页查询
func AlibabaMemberIsvPageQuery(clt *core.SDKClient, req *alimember.AlibabaMemberIsvPageQueryAPIRequest, session string) (*alimember.AlibabaMemberIsvPageQueryAPIResponse, error) {
	var resp alimember.AlibabaMemberIsvPageQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
