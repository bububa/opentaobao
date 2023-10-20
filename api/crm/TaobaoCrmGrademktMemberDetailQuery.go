package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmgrademktmemberdetailquery 会员等级营销-等级营销活动查询
// taobao.crm.grademkt.member.detail.query
//
// 商家通过该接口查询等级营销活动
func Taobaocrmgrademktmemberdetailquery(clt *core.SDKClient, req *crm.TaobaocrmgrademktmemberdetailqueryAPIRequest, session string) (*crm.TaobaocrmgrademktmemberdetailqueryAPIResponse, error) {
	var resp crm.TaobaocrmgrademktmemberdetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
