package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicaldepartmentsync 阿里健康预约挂号科室同步接口
// alibaba.alihealth.medical.department.sync
//
// 阿里健康预约挂号科室同步接口
func Alibabaalihealthmedicaldepartmentsync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicaldepartmentsyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicaldepartmentsyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicaldepartmentsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
