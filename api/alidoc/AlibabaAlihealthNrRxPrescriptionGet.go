package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// Alibabaalihealthnrrxprescriptionget 搜索处方详情
// alibaba.alihealth.nr.rx.prescription.get
//
// 获取互联网医院处方详情
func Alibabaalihealthnrrxprescriptionget(clt *core.SDKClient, req *alidoc.AlibabaalihealthnrrxprescriptiongetAPIRequest, session string) (*alidoc.AlibabaalihealthnrrxprescriptiongetAPIResponse, error) {
	var resp alidoc.AlibabaalihealthnrrxprescriptiongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
