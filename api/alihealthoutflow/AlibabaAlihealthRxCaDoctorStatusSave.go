package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthrxcadoctorstatussave ca认证获取医师认证结果
// alibaba.alihealth.rx.ca.doctor.status.save
//
// ca认证获取医师认证结果
func Alibabaalihealthrxcadoctorstatussave(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthrxcadoctorstatussaveAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthrxcadoctorstatussaveAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthrxcadoctorstatussaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
