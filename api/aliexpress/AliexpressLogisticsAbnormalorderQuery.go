package aliexpress

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpress"
)

// AliexpressLogisticsAbnormalorderQuery 异常订单查询
// aliexpress.logistics.abnormalorder.query
//
// 异常订单查询
func AliexpressLogisticsAbnormalorderQuery(clt *core.SDKClient, req *aliexpress.AliexpressLogisticsAbnormalorderQueryAPIRequest, resp *aliexpress.AliexpressLogisticsAbnormalorderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
