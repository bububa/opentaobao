package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowprescriptionupdate 处方外流-修改处方
// alibaba.alihealth.outflow.prescription.update
//
// 阿里健康-处方外流-对外提供处方修改功能
func Alibabaalihealthoutflowprescriptionupdate(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowprescriptionupdateAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowprescriptionupdateAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowprescriptionupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
