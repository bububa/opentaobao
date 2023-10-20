package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// Alibabaalihealthpwspecialsynchrosms 同步短信信息至阿里健康
// alibaba.alihealth.pw.special.synchrosms
//
// 同步短信信息至阿里健康
func Alibabaalihealthpwspecialsynchrosms(clt *core.SDKClient, req *alihealthpw.AlibabaalihealthpwspecialsynchrosmsAPIRequest, session string) (*alihealthpw.AlibabaalihealthpwspecialsynchrosmsAPIResponse, error) {
	var resp alihealthpw.AlibabaalihealthpwspecialsynchrosmsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
