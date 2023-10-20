package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolstairadditem 商品池阶梯商品添加
// alibaba.hm.marketing.itempool.stair.additem
//
// 添加商品池阶梯商品
func Alibabahmmarketingitempoolstairadditem(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolstairadditemAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolstairadditemAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolstairadditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
