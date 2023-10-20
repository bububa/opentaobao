package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Taobaowlbimportswaybillget 获取运单信息【云打印】
// taobao.wlb.imports.waybill.get
//
// 一般进口商家，获取订单的电子面单链接地址
func Taobaowlbimportswaybillget(clt *core.SDKClient, req *wlbimports.TaobaowlbimportswaybillgetAPIRequest, session string) (*wlbimports.TaobaowlbimportswaybillgetAPIResponse, error) {
	var resp wlbimports.TaobaowlbimportswaybillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
