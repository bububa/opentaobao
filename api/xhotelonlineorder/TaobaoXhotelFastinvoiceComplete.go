package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelfastinvoicecomplete 极速开票开票请求完成
// taobao.xhotel.fastinvoice.complete
//
// 极速开票开票请求回传,用于更新航信开票请求数据
func Taobaoxhotelfastinvoicecomplete(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelfastinvoicecompleteAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelfastinvoicecompleteAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelfastinvoicecompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
