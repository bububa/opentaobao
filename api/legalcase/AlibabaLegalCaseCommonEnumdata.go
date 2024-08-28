package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseCommonEnumdata 获取通用枚举值接口
// alibaba.legal.case.common.enumdata
//
// 获取通用枚举值接口
func AlibabaLegalCaseCommonEnumdata(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCommonEnumdataAPIRequest, resp *legalcase.AlibabaLegalCaseCommonEnumdataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
