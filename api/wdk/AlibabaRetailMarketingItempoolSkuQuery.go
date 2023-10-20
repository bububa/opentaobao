package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitempoolskuquery 查询商品池活动商品【同城零售】
// alibaba.retail.marketing.itempool.sku.query
//
// 查询商品池活动商品【同城零售】
func Alibabaretailmarketingitempoolskuquery(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitempoolskuqueryAPIRequest, session string) (*wdk.AlibabaretailmarketingitempoolskuqueryAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitempoolskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
