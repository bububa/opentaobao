package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelcommoninvoicelistvtwo 用户常用发票信息查询接口
// taobao.xhotel.commoninvoice.list.vtwo
//
// 获取用户常用发票信息接口
func Taobaoxhotelcommoninvoicelistvtwo(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelcommoninvoicelistvtwoAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelcommoninvoicelistvtwoAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelcommoninvoicelistvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
