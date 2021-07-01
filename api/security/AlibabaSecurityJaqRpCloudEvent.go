package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

/* AlibabaSecurityJaqRpCloudEvent
事件上报
alibaba.security.jaq.rp.cloud.event

事件上报接口 */
func AlibabaSecurityJaqRpCloudEvent(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudEventAPIRequest, session string) (*security.AlibabaSecurityJaqRpCloudEventAPIResponse, error) {
	var resp security.AlibabaSecurityJaqRpCloudEventAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
