package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// Tmallcrmmemberfrontunbind 品牌会员解绑
// tmall.crm.member.front.unbind
//
// 品牌会员解绑功能
func Tmallcrmmemberfrontunbind(clt *core.SDKClient, req *mei.TmallcrmmemberfrontunbindAPIRequest, session string) (*mei.TmallcrmmemberfrontunbindAPIResponse, error) {
	var resp mei.TmallcrmmemberfrontunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
