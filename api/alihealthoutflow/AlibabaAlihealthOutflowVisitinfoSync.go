package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowvisitinfosync 处方外流-问诊、处方同步
// alibaba.alihealth.outflow.visitinfo.sync
//
// 阿里健康-处方外流-对外提供问诊、处方功能
func Alibabaalihealthoutflowvisitinfosync(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowvisitinfosyncAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowvisitinfosyncAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowvisitinfosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
