package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

/* AlibabaRetailCommissionResultQuery
分佣结果查询
alibaba.retail.commission.result.query

查询导购分佣记录 */
func AlibabaRetailCommissionResultQuery(clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionResultQueryAPIRequest, session string) (*omniorder.AlibabaRetailCommissionResultQueryAPIResponse, error) {
	var resp omniorder.AlibabaRetailCommissionResultQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
