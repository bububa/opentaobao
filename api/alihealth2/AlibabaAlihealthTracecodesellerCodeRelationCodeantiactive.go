package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerCodeRelationCodeantiactive 解除码的关联关系
// alibaba.alihealth.tracecodeseller.code.relation.codeantiactive
//
// 解除码的关联关系
func AlibabaAlihealthTracecodesellerCodeRelationCodeantiactive(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthTracecodesellerCodeRelationCodeantiactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
