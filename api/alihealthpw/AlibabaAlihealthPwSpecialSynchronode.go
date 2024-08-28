package alihealthpw

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwSpecialSynchronode 合作方同步状态至阿里健康
// alibaba.alihealth.pw.special.synchronode
//
// 合作方同步状态至阿里健康
func AlibabaAlihealthPwSpecialSynchronode(ctx context.Context, clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwSpecialSynchronodeAPIRequest, resp *alihealthpw.AlibabaAlihealthPwSpecialSynchronodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
