package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelCommoninvoiceUpdate 常用发票信息更新接口
// taobao.xhotel.commoninvoice.update
//
// 常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加)
func TaobaoXhotelCommoninvoiceUpdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelCommoninvoiceUpdateAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelCommoninvoiceUpdateAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelCommoninvoiceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
