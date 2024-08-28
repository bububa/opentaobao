package alidoc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthNrRxPrescriptionGet 搜索处方详情
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
func AlibabaAlihealthNrRxPrescriptionGet(ctx context.Context, clt *core.SDKClient, req *alidoc.AlibabaAlihealthNrRxPrescriptionGetAPIRequest, resp *alidoc.AlibabaAlihealthNrRxPrescriptionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
