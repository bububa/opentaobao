package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamemberisvpagequery isv离线会员数据分页查询
// alibaba.member.isv.page.query
//
// isv离线会员数据分页查询
func Alibabamemberisvpagequery(clt *core.SDKClient, req *alimember.AlibabamemberisvpagequeryAPIRequest, session string) (*alimember.AlibabamemberisvpagequeryAPIResponse, error) {
	var resp alimember.AlibabamemberisvpagequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
