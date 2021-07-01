package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkMerchantStoreitemUpdate
修改门店商品
alibaba.wdk.merchant.storeitem.update

修改门店商品 */
func AlibabaWdkMerchantStoreitemUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantStoreitemUpdateAPIRequest, session string) (*wdk.AlibabaWdkMerchantStoreitemUpdateAPIResponse, error) {
	var resp wdk.AlibabaWdkMerchantStoreitemUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
