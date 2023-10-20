package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelcommoninvoiceupdate 常用发票信息更新接口
// taobao.xhotel.commoninvoice.update
//
// 常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加)
func Taobaoxhotelcommoninvoiceupdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelcommoninvoiceupdateAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelcommoninvoiceupdateAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelcommoninvoiceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
