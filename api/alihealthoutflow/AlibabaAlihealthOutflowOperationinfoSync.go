package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowoperationinfosync 处方外流-操作信息同步
// alibaba.alihealth.outflow.operationinfo.sync
//
// 阿里健康-处方外流-对外提供同步操作信息功能
func Alibabaalihealthoutflowoperationinfosync(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowoperationinfosyncAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowoperationinfosyncAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowoperationinfosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
