package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieThirdTelecomAutoauth 电信iot自动授权
// alibaba.ailabs.tmallgenie.third.telecom.autoauth
//
// 电信iot自动授权
func AlibabaAilabsTmallgenieThirdTelecomAutoauth(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest, session string) (*tmallgenie.AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
