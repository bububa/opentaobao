package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthDentalStoreInvisibleConsumeUpdate 门店无隐形消费签约
// alibaba.alihealth.dental.store.invisible.consume.update
//
// 门店无隐形消费签约
func AlibabaAlihealthDentalStoreInvisibleConsumeUpdate(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIRequest, session string) (*alihealth2.AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthDentalStoreInvisibleConsumeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
