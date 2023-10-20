package alihealthmedical

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthmedical"
)

// Alibabaalihealthmedicaldoctormsgsend 三方医生消息写入
// alibaba.alihealth.medical.doctor.msg.send
//
// 三方机构医生端发送消息同步写入阿里健康IM
func Alibabaalihealthmedicaldoctormsgsend(clt *core.SDKClient, req *alihealthmedical.AlibabaalihealthmedicaldoctormsgsendAPIRequest, session string) (*alihealthmedical.AlibabaalihealthmedicaldoctormsgsendAPIResponse, error) {
	var resp alihealthmedical.AlibabaalihealthmedicaldoctormsgsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
