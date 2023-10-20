package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmmembersget 获取卖家的会员（基本查询）
// taobao.crm.members.get
//
// 查询卖家的会员，进行基本的查询，返回符合条件的会员列表
func Taobaocrmmembersget(clt *core.SDKClient, req *crm.TaobaocrmmembersgetAPIRequest, session string) (*crm.TaobaocrmmembersgetAPIResponse, error) {
	var resp crm.TaobaocrmmembersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
