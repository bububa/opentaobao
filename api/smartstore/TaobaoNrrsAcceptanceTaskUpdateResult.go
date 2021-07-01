package smartstore

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/smartstore"
)

/* TaobaoNrrsAcceptanceTaskUpdateResult
更新验收任务支付宝结果
taobao.nrrs.acceptance.task.updateResult

智慧门店商家验收任务检查相关接口-更新支付宝的验收结果。 */
func TaobaoNrrsAcceptanceTaskUpdateResult(clt *core.SDKClient, req *smartstore.TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest, session string) (*smartstore.TaobaoNrrsAcceptanceTaskUpdateResultAPIResponse, error) {
	var resp smartstore.TaobaoNrrsAcceptanceTaskUpdateResultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
