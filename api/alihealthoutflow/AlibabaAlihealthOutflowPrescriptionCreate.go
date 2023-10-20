package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowprescriptioncreate 处方外流-创建处方
// alibaba.alihealth.outflow.prescription.create
//
// 阿里健康-处方外流-对外提供保存处方功能
func Alibabaalihealthoutflowprescriptioncreate(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowprescriptioncreateAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowprescriptioncreateAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowprescriptioncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
