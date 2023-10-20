package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicaldoctorsync 阿里健康预约挂号医生同步接口
// alibaba.alihealth.medical.doctor.sync
//
// 阿里健康预约挂号医生同步接口
func Alibabaalihealthmedicaldoctorsync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicaldoctorsyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicaldoctorsyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicaldoctorsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
