package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelMemberAlipayQuery 希尔顿会员查询
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
func TaobaoXhotelMemberAlipayQuery(clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelMemberAlipayQueryAPIRequest, resp *xhotelcrm.TaobaoXhotelMemberAlipayQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
