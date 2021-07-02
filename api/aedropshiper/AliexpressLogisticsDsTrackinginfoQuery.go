package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressLogisticsDsTrackinginfoQuery 查询物流追踪信息
// aliexpress.logistics.ds.trackinginfo.query
//
// Dropshipper查询物流追踪信息
func AliexpressLogisticsDsTrackinginfoQuery(clt *core.SDKClient, req *aedropshiper.AliexpressLogisticsDsTrackinginfoQueryAPIRequest, session string) (*aedropshiper.AliexpressLogisticsDsTrackinginfoQueryAPIResponse, error) {
	var resp aedropshiper.AliexpressLogisticsDsTrackinginfoQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
