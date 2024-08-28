package alihealthpw

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwSpecialSynchropatientname 同步患者姓名至阿里健康
// alibaba.alihealth.pw.special.synchropatientname
//
// 同步患者姓名至阿里健康
func AlibabaAlihealthPwSpecialSynchropatientname(ctx context.Context, clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwSpecialSynchropatientnameAPIRequest, resp *alihealthpw.AlibabaAlihealthPwSpecialSynchropatientnameAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
