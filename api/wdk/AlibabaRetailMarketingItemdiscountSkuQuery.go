package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaretailmarketingitemdiscountskuquery 查询单品特价活动商品【同城零售】
// alibaba.retail.marketing.itemdiscount.sku.query
//
// 查询单品特价活动商品【同城零售】
func Alibabaretailmarketingitemdiscountskuquery(clt *core.SDKClient, req *wdk.AlibabaretailmarketingitemdiscountskuqueryAPIRequest, session string) (*wdk.AlibabaretailmarketingitemdiscountskuqueryAPIResponse, error) {
	var resp wdk.AlibabaretailmarketingitemdiscountskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
