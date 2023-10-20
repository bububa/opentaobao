package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberinfoUpdatePrivy 编辑会员资料
// taobao.crm.memberinfo.update.privy
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
func TaobaoCrmMemberinfoUpdatePrivy(clt *core.SDKClient, req *crm.TaobaoCrmMemberinfoUpdatePrivyAPIRequest, session string) (*crm.TaobaoCrmMemberinfoUpdatePrivyAPIResponse, error) {
	var resp crm.TaobaoCrmMemberinfoUpdatePrivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
