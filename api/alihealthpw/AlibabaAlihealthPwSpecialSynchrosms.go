package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwSpecialSynchrosms 同步短信信息至阿里健康
// alibaba.alihealth.pw.special.synchrosms
//
// 同步短信信息至阿里健康
func AlibabaAlihealthPwSpecialSynchrosms(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwSpecialSynchrosmsAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwSpecialSynchrosmsAPIResponse, error) {
	var resp alihealthpw.AlibabaAlihealthPwSpecialSynchrosmsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
