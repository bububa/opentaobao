package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkSkuQuery
查询商品
alibaba.wdk.sku.query

查询商品 */
func AlibabaWdkSkuQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuQueryAPIRequest, session string) (*wdk.AlibabaWdkSkuQueryAPIResponse, error) {
	var resp wdk.AlibabaWdkSkuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
