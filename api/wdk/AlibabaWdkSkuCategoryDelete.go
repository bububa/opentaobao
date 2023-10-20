package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCategoryDelete 商家类目删除接口
// alibaba.wdk.sku.category.delete
//
// 商家类目删除接口
func AlibabaWdkSkuCategoryDelete(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryDeleteAPIRequest, session string) (*wdk.AlibabaWdkSkuCategoryDeleteAPIResponse, error) {
	var resp wdk.AlibabaWdkSkuCategoryDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
