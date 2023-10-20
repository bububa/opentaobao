package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

// TaobaoRdcAligeniusIdentificationCaseUpdate 鉴定工单信息同步
// taobao.rdc.aligenius.identification.case.update
//
// 同步商家鉴定工单信息
func TaobaoRdcAligeniusIdentificationCaseUpdate(clt *core.SDKClient, req *refund.TaobaoRdcAligeniusIdentificationCaseUpdateAPIRequest, resp *refund.TaobaoRdcAligeniusIdentificationCaseUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
