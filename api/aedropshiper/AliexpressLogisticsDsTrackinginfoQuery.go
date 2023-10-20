package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressLogisticsDsTrackinginfoQuery 查询物流追踪信息
// aliexpress.logistics.ds.trackinginfo.query
//
// Dropshipper查询物流追踪信息
func AliexpressLogisticsDsTrackinginfoQuery(clt *core.SDKClient, req *aedropshiper.AliexpressLogisticsDsTrackinginfoQueryAPIRequest, resp *aedropshiper.AliexpressLogisticsDsTrackinginfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
