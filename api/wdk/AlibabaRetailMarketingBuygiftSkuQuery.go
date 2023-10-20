package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingbuygiftskuquery 查询买赠活动商品【同城零售】
// alibaba.retail.marketing.buygift.sku.query
//
// 查询买赠活动商品【同城零售】
func Alibabaretailmarketingbuygiftskuquery(clt *core.SDKClient, req *wdk.AlibabaretailmarketingbuygiftskuqueryAPIRequest, session string) (*wdk.AlibabaretailmarketingbuygiftskuqueryAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingbuygiftskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
