package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberAlipayQuery 希尔顿会员查询
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
func TaobaoXhotelMemberAlipayQuery(clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberAlipayQueryAPIRequest, session string) (*xhotelcrm.TaobaoXhotelMemberAlipayQueryAPIResponse, error) {
	var resp xhotelcrm.TaobaoXhotelMemberAlipayQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
