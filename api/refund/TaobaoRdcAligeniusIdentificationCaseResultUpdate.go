package refund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/refund"
)

/* TaobaoRdcAligeniusIdentificationCaseResultUpdate
鉴定工单结果同步
taobao.rdc.aligenius.identification.case.result.update

同步鉴定工单结果信息 */
func TaobaoRdcAligeniusIdentificationCaseResultUpdate(clt *core.SDKClient, req *refund.TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIRequest, session string) (*refund.TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponse, error) {
	var resp refund.TaobaoRdcAligeniusIdentificationCaseResultUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
