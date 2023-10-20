package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkitemmorebarcodeops 商品一品多码维护操作
// alibaba.wdk.item.morebarcode.ops
//
// 商品一品多码维护操作
func Alibabawdkitemmorebarcodeops(clt *core.SDKClient, req *wdkitem.AlibabawdkitemmorebarcodeopsAPIRequest, session string) (*wdkitem.AlibabawdkitemmorebarcodeopsAPIResponse, error) {
	var resp wdkitem.AlibabawdkitemmorebarcodeopsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
