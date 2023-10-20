package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelCommoninvoiceRemove 常用发票信息删除接口
// taobao.xhotel.commoninvoice.remove
//
// 常用发票信息删除接口
func TaobaoXhotelCommoninvoiceRemove(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelCommoninvoiceRemoveAPIRequest, resp *xhotelonlineorder.TaobaoXhotelCommoninvoiceRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
