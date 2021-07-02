package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuUpdate 更新商品
// alibaba.wdk.sku.update
//
// 开放商品更新接口
func AlibabaWdkSkuUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkSkuUpdateAPIRequest, session string) (*wdk.AlibabaWdkSkuUpdateAPIResponse, error) {
	var resp wdk.AlibabaWdkSkuUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
