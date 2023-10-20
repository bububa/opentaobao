package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinFlowWalletGrade 获取流量档位
// alibaba.aliqin.flow.wallet.grade
//
// 获取直充流量档位
func AlibabaAliqinFlowWalletGrade(clt *core.SDKClient, req *alicom.AlibabaAliqinFlowWalletGradeAPIRequest, resp *alicom.AlibabaAliqinFlowWalletGradeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
