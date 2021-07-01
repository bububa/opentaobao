package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

/* AlibabaAlihealthDentalStoreInsertorupdate
ISV新增/修改口腔门店
alibaba.alihealth.dental.store.insertorupdate

ISV新增/修改口腔门店 */
func AlibabaAlihealthDentalStoreInsertorupdate(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStoreInsertorupdateAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalStoreInsertorupdateAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDentalStoreInsertorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
