package alihealthpw

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwSpecialSynchrosms 同步短信信息至阿里健康
// alibaba.alihealth.pw.special.synchrosms
//
// 同步短信信息至阿里健康
func AlibabaAlihealthPwSpecialSynchrosms(ctx context.Context, clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwSpecialSynchrosmsAPIRequest, resp *alihealthpw.AlibabaAlihealthPwSpecialSynchrosmsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
