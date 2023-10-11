package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolAddcategory 增加商品池里面的类目
// alibaba.hm.marketing.itempool.addcategory
//
// 增加商品池里面的类目
func AlibabaHmMarketingItempoolAddcategory(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolAddcategoryAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolAddcategoryAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolAddcategoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
