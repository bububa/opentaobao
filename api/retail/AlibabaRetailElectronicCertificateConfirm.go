package retail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// AlibabaRetailElectronicCertificateConfirm 确认核销接口
// alibaba.retail.electronic.certificate.confirm
//
// 确认核销接口
func AlibabaRetailElectronicCertificateConfirm(ctx context.Context, clt *core.SDKClient, req *retail.AlibabaRetailElectronicCertificateConfirmAPIRequest, resp *retail.AlibabaRetailElectronicCertificateConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
