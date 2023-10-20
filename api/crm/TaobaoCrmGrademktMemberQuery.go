package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmgrademktmemberquery 会员等级营销-会员关系查询
// taobao.crm.grademkt.member.query
//
// 商家通过该接口查询线上店铺会员。
func Taobaocrmgrademktmemberquery(clt *core.SDKClient, req *crm.TaobaocrmgrademktmemberqueryAPIRequest, session string) (*crm.TaobaocrmgrademktmemberqueryAPIResponse, error) {
	var resp crm.TaobaocrmgrademktmemberqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
