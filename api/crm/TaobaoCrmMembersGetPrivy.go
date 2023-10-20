package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmmembersgetprivy 获取卖家会员(基本查询)
// taobao.crm.members.get.privy
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
func Taobaocrmmembersgetprivy(clt *core.SDKClient, req *crm.TaobaocrmmembersgetprivyAPIRequest, session string) (*crm.TaobaocrmmembersgetprivyAPIResponse, error) {
	var resp crm.TaobaocrmmembersgetprivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
