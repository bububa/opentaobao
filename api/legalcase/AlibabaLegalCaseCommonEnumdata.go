package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseCommonEnumdata 获取通用枚举值接口
// alibaba.legal.case.common.enumdata
//
// 获取通用枚举值接口
func AlibabaLegalCaseCommonEnumdata(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCommonEnumdataAPIRequest, resp *legalcase.AlibabaLegalCaseCommonEnumdataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
