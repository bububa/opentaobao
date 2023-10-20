package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtmemberopenid 根据会员手机查询openId
// tmall.nrt.member.openid
//
// 根据会员手机查询openId
func Tmallnrtmemberopenid(clt *core.SDKClient, req *nrt.TmallnrtmemberopenidAPIRequest, session string) (*nrt.TmallnrtmemberopenidAPIResponse, error) {
	var resp nrt.TmallnrtmemberopenidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
