package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthNrRxPrescriptionGet 搜索处方详情
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
func AlibabaAlihealthNrRxPrescriptionGet(clt *core.SDKClient, req *alidoc.AlibabaAlihealthNrRxPrescriptionGetAPIRequest, session string) (*alidoc.AlibabaAlihealthNrRxPrescriptionGetAPIResponse, error) {
	var resp alidoc.AlibabaAlihealthNrRxPrescriptionGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
