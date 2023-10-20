package alitripcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripcar"
)

// Alitripcarinvoiceconfirm 发票确认接口
// alitrip.car.invoice.confirm
//
// 飞猪发票回调接口
func Alitripcarinvoiceconfirm(clt *core.SDKClient, req *alitripcar.AlitripcarinvoiceconfirmAPIRequest, session string) (*alitripcar.AlitripcarinvoiceconfirmAPIResponse, error) {
	var resp alitripcar.AlitripcarinvoiceconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
