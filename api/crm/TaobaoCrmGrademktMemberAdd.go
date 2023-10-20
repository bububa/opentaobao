package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmgrademktmemberadd 会员等级营销-会员吸纳
// taobao.crm.grademkt.member.add
//
// 商家通过该接口吸纳线上店铺会员。
func Taobaocrmgrademktmemberadd(clt *core.SDKClient, req *crm.TaobaocrmgrademktmemberaddAPIRequest, session string) (*crm.TaobaocrmgrademktmemberaddAPIResponse, error) {
	var resp crm.TaobaocrmgrademktmemberaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
