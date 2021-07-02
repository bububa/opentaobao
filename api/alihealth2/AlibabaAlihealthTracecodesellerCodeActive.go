package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerCodeActive 上传激活码的文件
// alibaba.alihealth.tracecodeseller.code.active
//
// 上传商品的激活码文件，存到系统中
func AlibabaAlihealthTracecodesellerCodeActive(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerCodeActiveAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerCodeActiveAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthTracecodesellerCodeActiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
