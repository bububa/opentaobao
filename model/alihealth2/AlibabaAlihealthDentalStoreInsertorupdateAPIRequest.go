package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalStoreInsertorupdateAPIRequest
ISV新增/修改口腔门店 API请求
alibaba.alihealth.dental.store.insertorupdate

ISV新增/修改口腔门店 */
type AlibabaAlihealthDentalStoreInsertorupdateAPIRequest struct {
	model.Params
	// 门店
	_store *DentalOuterStoreRequest
}

// New
