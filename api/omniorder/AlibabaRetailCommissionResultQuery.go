package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Alibabaretailcommissionresultquery 分佣结果查询
// alibaba.retail.commission.result.query
//
// 查询导购分佣记录
func Alibabaretailcommissionresultquery(clt *core.SDKClient, req *omniorder.AlibabaretailcommissionresultqueryAPIRequest, session string) (*omniorder.AlibabaretailcommissionresultqueryAPIResponse, error) {
	var resp omniorder.AlibabaretailcommissionresultqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
