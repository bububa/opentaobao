package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBcItemPeriodSync 代销品效期同步
// alibaba.alihealth.bc.item.period.sync
//
// 代销品效期同步
func AlibabaAlihealthBcItemPeriodSync(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBcItemPeriodSyncAPIRequest, session string) (*alihealth2.AlibabaAlihealthBcItemPeriodSyncAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthBcItemPeriodSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
