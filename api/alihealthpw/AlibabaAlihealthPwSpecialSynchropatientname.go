package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// Alibabaalihealthpwspecialsynchropatientname 同步患者姓名至阿里健康
// alibaba.alihealth.pw.special.synchropatientname
//
// 同步患者姓名至阿里健康
func Alibabaalihealthpwspecialsynchropatientname(clt *core.SDKClient, req *alihealthpw.AlibabaalihealthpwspecialsynchropatientnameAPIRequest, session string) (*alihealthpw.AlibabaalihealthpwspecialsynchropatientnameAPIResponse, error) {
	var resp alihealthpw.AlibabaalihealthpwspecialsynchropatientnameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
