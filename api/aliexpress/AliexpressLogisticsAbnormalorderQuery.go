package aliexpress

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpress"
)

// Aliexpresslogisticsabnormalorderquery 异常订单查询
// aliexpress.logistics.abnormalorder.query
//
// 异常订单查询
func Aliexpresslogisticsabnormalorderquery(clt *core.SDKClient, req *aliexpress.AliexpresslogisticsabnormalorderqueryAPIRequest, session string) (*aliexpress.AliexpresslogisticsabnormalorderqueryAPIResponse, error) {
	var resp aliexpress.AliexpresslogisticsabnormalorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
