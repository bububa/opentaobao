package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempooladdcategory 增加商品池里面的类目
// alibaba.wdk.marketing.itempool.addcategory
//
// 增加商品池里面的类目
func Alibabawdkmarketingitempooladdcategory(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempooladdcategoryAPIRequest, session string) (*wdk.AlibabawdkmarketingitempooladdcategoryAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempooladdcategoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
