package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// AlibabaRetailElectronicCertificatePreConfirm 贩卖机开始核销接口
// alibaba.retail.electronic.certificate.pre.confirm
//
// 零售终端贩卖机开始核销接口,返回待领的商品ID
func AlibabaRetailElectronicCertificatePreConfirm(clt *core.SDKClient, req *retail.AlibabaRetailElectronicCertificatePreConfirmAPIRequest, session string) (*retail.AlibabaRetailElectronicCertificatePreConfirmAPIResponse, error) {
	var resp retail.AlibabaRetailElectronicCertificatePreConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
