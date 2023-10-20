package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// Taobaoxhotelmemberalipayquery 希尔顿会员查询
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
func Taobaoxhotelmemberalipayquery(clt *core.SDKClient, req *xhotelcrm.TaobaoxhotelmemberalipayqueryAPIRequest, session string) (*xhotelcrm.TaobaoxhotelmemberalipayqueryAPIResponse, error) {
	var resp xhotelcrm.TaobaoxhotelmemberalipayqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
