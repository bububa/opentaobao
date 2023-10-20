package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodeplatformCodeActive 正大鸡蛋激活追溯码
// alibaba.alihealth.tracecodeplatform.code.active
//
// 用于正大鸡蛋激活追溯码
func AlibabaAlihealthTracecodeplatformCodeActive(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodeplatformCodeActiveAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
