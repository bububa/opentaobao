package aecreatives

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// Aliexpressaffiliatehotproductquery 查询联盟爆品数据
// aliexpress.affiliate.hotproduct.query
//
// 查询联盟爆品API
func Aliexpressaffiliatehotproductquery(clt *core.SDKClient, req *aecreatives.AliexpressaffiliatehotproductqueryAPIRequest, session string) (*aecreatives.AliexpressaffiliatehotproductqueryAPIResponse, error) {
	var resp aecreatives.AliexpressaffiliatehotproductqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
