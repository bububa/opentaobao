package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieThirdTelecomPushrender 电信-推送拉起设备应用
// alibaba.ailabs.tmallgenie.third.telecom.pushrender
//
// 电信-推送拉起设备应用
func AlibabaAilabsTmallgenieThirdTelecomPushrender(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIRequest, session string) (*tmallgenie.AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
