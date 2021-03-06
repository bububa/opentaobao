package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelCommoninvoiceRemove 常用发票信息删除接口
// taobao.xhotel.commoninvoice.remove
//
// 常用发票信息删除接口
func TaobaoXhotelCommoninvoiceRemove(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelCommoninvoiceRemoveAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelCommoninvoiceRemoveAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelCommoninvoiceRemoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
