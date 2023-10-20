package aliexpress

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpress"
)

// AliexpressLogisticsAbnormalorderQuery 异常订单查询
// aliexpress.logistics.abnormalorder.query
//
// 异常订单查询
func AliexpressLogisticsAbnormalorderQuery(clt *core.SDKClient, req *aliexpress.AliexpressLogisticsAbnormalorderQueryAPIRequest, session string) (*aliexpress.AliexpressLogisticsAbnormalorderQueryAPIResponse, error) {
	var resp aliexpress.AliexpressLogisticsAbnormalorderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
