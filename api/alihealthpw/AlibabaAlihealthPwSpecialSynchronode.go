package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// Alibabaalihealthpwspecialsynchronode 合作方同步状态至阿里健康
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
func Alibabaalihealthpwspecialsynchronode(clt *core.SDKClient, req *alihealthpw.AlibabaalihealthpwspecialsynchronodeAPIRequest, session string) (*alihealthpw.AlibabaalihealthpwspecialsynchronodeAPIResponse, error) {
	var resp alihealthpw.AlibabaalihealthpwspecialsynchronodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
