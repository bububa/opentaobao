package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowpatientinfosync 处方外流-患者基础信息同步
// alibaba.alihealth.outflow.patientinfo.sync
//
// 阿里健康-处方外流-对外提供同步患者基础信息功能
func Alibabaalihealthoutflowpatientinfosync(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowpatientinfosyncAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowpatientinfosyncAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowpatientinfosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
