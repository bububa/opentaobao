package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelpmsguestbillgetvtwo 客人PMS账单信息查询
// taobao.xhotel.pms.guestbill.get.vtwo
//
// 从pms获取客人账单信息
func Taobaoxhotelpmsguestbillgetvtwo(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelpmsguestbillgetvtwoAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelpmsguestbillgetvtwoAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelpmsguestbillgetvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
