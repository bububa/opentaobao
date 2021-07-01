package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthDentalStoreAuditQuery
ISV查询门店审核状态
alibaba.alihealth.dental.store.audit.query

ISV查询门店审核状态 */
func AlibabaAlihealthDentalStoreAuditQuery(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStoreAuditQueryAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalStoreAuditQueryAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDentalStoreAuditQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
