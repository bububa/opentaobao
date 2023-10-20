package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmgrademktmemberdetailcreate 会员等级营销-创建商品等级营销明细
// taobao.crm.grademkt.member.detail.create
//
// 创建商品等级营销明细
func Taobaocrmgrademktmemberdetailcreate(clt *core.SDKClient, req *crm.TaobaocrmgrademktmemberdetailcreateAPIRequest, session string) (*crm.TaobaocrmgrademktmemberdetailcreateAPIResponse, error) {
	var resp crm.TaobaocrmgrademktmemberdetailcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
