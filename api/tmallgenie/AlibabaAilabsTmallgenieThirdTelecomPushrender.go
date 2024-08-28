package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsTmallgenieThirdTelecomPushrender 电信-推送拉起设备应用
// alibaba.ailabs.tmallgenie.third.telecom.pushrender
//
// 电信-推送拉起设备应用
func AlibabaAilabsTmallgenieThirdTelecomPushrender(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIRequest, resp *tmallgenie.AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
