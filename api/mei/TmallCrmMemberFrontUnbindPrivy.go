package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// Tmallcrmmemberfrontunbindprivy 品牌会员解绑(隐私号版）
// tmall.crm.member.front.unbind.privy
//
// 品牌会员解绑功能
func Tmallcrmmemberfrontunbindprivy(clt *core.SDKClient, req *mei.TmallcrmmemberfrontunbindprivyAPIRequest, session string) (*mei.TmallcrmmemberfrontunbindprivyAPIResponse, error) {
	var resp mei.TmallcrmmemberfrontunbindprivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
