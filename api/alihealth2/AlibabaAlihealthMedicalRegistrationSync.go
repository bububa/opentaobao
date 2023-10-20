package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalregistrationsync 阿里健康支付宝挂号记录回传接口
// alibaba.alihealth.medical.registration.sync
//
// 阿里健康支付宝挂号记录回传接口
func Alibabaalihealthmedicalregistrationsync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalregistrationsyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalregistrationsyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalregistrationsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
