package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalStoreInsertorupdate ISV新增/修改口腔门店
// alibaba.alihealth.dental.store.insertorupdate
//
// ISV新增/修改口腔门店
func AlibabaAlihealthDentalStoreInsertorupdate(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStoreInsertorupdateAPIRequest, resp *alihealth2.AlibabaAlihealthDentalStoreInsertorupdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
