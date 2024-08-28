package alihealthpw

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwApplynodeUpdate 申请节点变更回调
// alibaba.alihealth.pw.applynode.update
//
// 基金会回调阿里健康更新药品援助申请单的状态
func AlibabaAlihealthPwApplynodeUpdate(ctx context.Context, clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwApplynodeUpdateAPIRequest, resp *alihealthpw.AlibabaAlihealthPwApplynodeUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
