package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmmemberinfoupdateprivy 编辑会员资料
// taobao.crm.memberinfo.update.privy
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
func Taobaocrmmemberinfoupdateprivy(clt *core.SDKClient, req *crm.TaobaocrmmemberinfoupdateprivyAPIRequest, session string) (*crm.TaobaocrmmemberinfoupdateprivyAPIResponse, error) {
	var resp crm.TaobaocrmmemberinfoupdateprivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
