package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthNrRxPrescriptionGet 搜索处方详情
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
func AlibabaAlihealthNrRxPrescriptionGet(clt *core.SDKClient, req *alidoc.AlibabaAlihealthNrRxPrescriptionGetAPIRequest, resp *alidoc.AlibabaAlihealthNrRxPrescriptionGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
