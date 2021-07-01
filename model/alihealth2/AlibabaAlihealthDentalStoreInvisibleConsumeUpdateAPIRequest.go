package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest
门店无隐形消费签约 API请求
alibaba.alihealth.dental.store.invisible.consume.update

门店无隐形消费签约 */
type AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest struct {
	model.Params
	// 入参
	_store *DentalOuterStoreNicRequest
}

// New
