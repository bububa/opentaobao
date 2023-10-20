package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelPmsGuestbillGetVtwo 客人PMS账单信息查询
// taobao.xhotel.pms.guestbill.get.vtwo
//
// 从pms获取客人账单信息
func TaobaoXhotelPmsGuestbillGetVtwo(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelPmsGuestbillGetVtwoAPIRequest, resp *xhotelonlineorder.TaobaoXhotelPmsGuestbillGetVtwoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
