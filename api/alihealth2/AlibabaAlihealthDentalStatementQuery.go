package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthDentalStatementQuery
ISV查询对账单
alibaba.alihealth.dental.statement.query

ISV查询对账单 */
func AlibabaAlihealthDentalStatementQuery(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStatementQueryAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalStatementQueryAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDentalStatementQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
