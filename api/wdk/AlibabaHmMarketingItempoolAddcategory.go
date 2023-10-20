package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempooladdcategory 增加商品池里面的类目
// alibaba.hm.marketing.itempool.addcategory
//
// 增加商品池里面的类目
func Alibabahmmarketingitempooladdcategory(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempooladdcategoryAPIRequest, session string) (*wdk.AlibabahmmarketingitempooladdcategoryAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempooladdcategoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
