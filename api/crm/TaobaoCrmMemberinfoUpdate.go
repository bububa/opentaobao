package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberinfoUpdate 编辑会员资料
// taobao.crm.memberinfo.update
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
func TaobaoCrmMemberinfoUpdate(clt *core.SDKClient, req *crm.TaobaoCrmMemberinfoUpdateAPIRequest, resp *crm.TaobaoCrmMemberinfoUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
