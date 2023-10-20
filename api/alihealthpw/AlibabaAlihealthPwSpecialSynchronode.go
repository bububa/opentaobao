package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwSpecialSynchronode 合作方同步状态至阿里健康
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
func AlibabaAlihealthPwSpecialSynchronode(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwSpecialSynchronodeAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwSpecialSynchronodeAPIResponse, error) {
	var resp alihealthpw.AlibabaAlihealthPwSpecialSynchronodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
