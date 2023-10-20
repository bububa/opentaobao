package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelcommoninvoiceremove 常用发票信息删除接口
// taobao.xhotel.commoninvoice.remove
//
// 常用发票信息删除接口
func Taobaoxhotelcommoninvoiceremove(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelcommoninvoiceremoveAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelcommoninvoiceremoveAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelcommoninvoiceremoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
