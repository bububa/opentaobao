package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabapricepromotionitemadd 新增档期商品
// alibaba.price.promotion.item.add
//
// 批量新增档期活动商品
func Alibabapricepromotionitemadd(clt *core.SDKClient, req *wdk.AlibabapricepromotionitemaddAPIRequest, session string) (*wdk.AlibabapricepromotionitemaddAPIResponse, error) {
	var resp wdk.AlibabapricepromotionitemaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
