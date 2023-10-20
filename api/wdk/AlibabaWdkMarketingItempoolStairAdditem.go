package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolstairadditem 商品池阶梯商品添加
// alibaba.wdk.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
func Alibabawdkmarketingitempoolstairadditem(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolstairadditemAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolstairadditemAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolstairadditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
