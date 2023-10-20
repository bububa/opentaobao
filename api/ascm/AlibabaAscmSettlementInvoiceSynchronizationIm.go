package ascm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascm"
)

// Alibabaascmsettlementinvoicesynchronizationim 英迈发票同步到结算
// alibaba.ascm.settlement.invoice.synchronization.im
//
// 外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
func Alibabaascmsettlementinvoicesynchronizationim(clt *core.SDKClient, req *ascm.AlibabaascmsettlementinvoicesynchronizationimAPIRequest, session string) (*ascm.AlibabaascmsettlementinvoicesynchronizationimAPIResponse, error) {
	var resp ascm.AlibabaascmsettlementinvoicesynchronizationimAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
