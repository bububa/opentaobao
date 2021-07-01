package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkMerchantItemCreatedraft
新建商品草稿
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口 */
func AlibabaWdkMerchantItemCreatedraft(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantItemCreatedraftAPIRequest, session string) (*wdk.AlibabaWdkMerchantItemCreatedraftAPIResponse, error) {
	var resp wdk.AlibabaWdkMerchantItemCreatedraftAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
