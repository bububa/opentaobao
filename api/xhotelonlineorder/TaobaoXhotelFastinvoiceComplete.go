package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* TaobaoXhotelFastinvoiceComplete
极速开票开票请求完成
taobao.xhotel.fastinvoice.complete

极速开票开票请求回传,用于更新航信开票请求数据 */
func TaobaoXhotelFastinvoiceComplete(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelFastinvoiceCompleteAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelFastinvoiceCompleteAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelFastinvoiceCompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
