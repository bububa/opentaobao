package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelCommoninvoiceListVtwo 用户常用发票信息查询接口
// taobao.xhotel.commoninvoice.list.vtwo
//
// 获取用户常用发票信息接口
func TaobaoXhotelCommoninvoiceListVtwo(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelCommoninvoiceListVtwoAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelCommoninvoiceListVtwoAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelCommoninvoiceListVtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
