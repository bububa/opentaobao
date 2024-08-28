package alihealthpw

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwApplynodeUpdatename 回调变更患者姓名
// alibaba.alihealth.pw.applynode.updatename
//
// 回调变更患者姓名
func AlibabaAlihealthPwApplynodeUpdatename(ctx context.Context, clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwApplynodeUpdatenameAPIRequest, resp *alihealthpw.AlibabaAlihealthPwApplynodeUpdatenameAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
