package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoAdd 自促活动申请接口
// taobao.xhotel.bnbpromo.add
//
// 自促活动申请接口
func TaobaoXhotelBnbpromoAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
