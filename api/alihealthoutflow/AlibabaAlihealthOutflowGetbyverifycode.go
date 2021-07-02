package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthOutflowGetbyverifycode 处方外流药店通过核销码获取处方
// alibaba.alihealth.outflow.getbyverifycode
//
// 阿里健康对合作药店提供通过核销码查看处方的功能
func AlibabaAlihealthOutflowGetbyverifycode(clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthOutflowGetbyverifycodeAPIRequest, session string) (*alihealthoutflow.AlibabaAlihealthOutflowGetbyverifycodeAPIResponse, error) {
	var resp alihealthoutflow.AlibabaAlihealthOutflowGetbyverifycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
