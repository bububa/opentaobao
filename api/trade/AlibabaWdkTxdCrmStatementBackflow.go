package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* AlibabaWdkTxdCrmStatementBackflow
淘鲜达商家会员账单回流
alibaba.wdk.txd.crm.statement.backflow

淘鲜达商家会员账单回流 */
func AlibabaWdkTxdCrmStatementBackflow(clt *core.SDKClient, req *trade.AlibabaWdkTxdCrmStatementBackflowAPIRequest, session string) (*trade.AlibabaWdkTxdCrmStatementBackflowAPIResponse, error) {
	var resp trade.AlibabaWdkTxdCrmStatementBackflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
