package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// Alibabaretailelectroniccertificatepreconfirm 贩卖机开始核销接口
// alibaba.retail.electronic.certificate.pre.confirm
//
// 零售终端贩卖机开始核销接口,返回待领的商品ID
func Alibabaretailelectroniccertificatepreconfirm(clt *core.SDKClient, req *retail.AlibabaretailelectroniccertificatepreconfirmAPIRequest, session string) (*retail.AlibabaretailelectroniccertificatepreconfirmAPIResponse, error) {
	var resp retail.AlibabaretailelectroniccertificatepreconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
