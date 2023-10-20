package alihealth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth"
)

// Alibabaalihealthprescriptionauthget 阿里健康处方平台获取授权码
// alibaba.alihealth.prescription.auth.get
//
// 获取处方用户授权
func Alibabaalihealthprescriptionauthget(clt *core.SDKClient, req *alihealth.AlibabaalihealthprescriptionauthgetAPIRequest, session string) (*alihealth.AlibabaalihealthprescriptionauthgetAPIResponse, error) {
	var resp alihealth.AlibabaalihealthprescriptionauthgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
