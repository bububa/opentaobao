package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletGrade 获取流量档位
// alibaba.aliqin.flow.wallet.grade
//
// 获取直充流量档位
func AlibabaAliqinFlowWalletGrade(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletGradeAPIRequest, resp *alicom.AlibabaAliqinFlowWalletGradeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
