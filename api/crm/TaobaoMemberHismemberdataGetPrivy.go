package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoMemberHismemberdataGetPrivy 会员历史备份数据查询
// taobao.member.hismemberdata.get.privy
//
// 会员历史备份数据分页查询，查询内容为等级，会员备份版本，会员nick等信息.
func TaobaoMemberHismemberdataGetPrivy(clt *core.SDKClient, req *crm.TaobaoMemberHismemberdataGetPrivyAPIRequest, session string) (*crm.TaobaoMemberHismemberdataGetPrivyAPIResponse, error) {
	var resp crm.TaobaoMemberHismemberdataGetPrivyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
