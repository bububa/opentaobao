package legalsuit

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitCaseGet 获取案件信息接口v2版本
// alibaba.legal.suit.case.get
//
// 获取案件信息
func AlibabaLegalSuitCaseGet(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCaseGetAPIRequest, resp *legalsuit.AlibabaLegalSuitCaseGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
