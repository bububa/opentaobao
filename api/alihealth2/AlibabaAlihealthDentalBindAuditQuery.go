package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthDentalBindAuditQuery
ISV查询绑定审核状态
alibaba.alihealth.dental.bind.audit.query

ISV查询绑定审核状态 */
func AlibabaAlihealthDentalBindAuditQuery(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalBindAuditQueryAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalBindAuditQueryAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDentalBindAuditQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
