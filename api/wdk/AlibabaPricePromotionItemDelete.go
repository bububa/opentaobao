package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabapricepromotionitemdelete 批量删除档期
// alibaba.price.promotion.item.delete
//
// 盒马帮批量删除档期商品
func Alibabapricepromotionitemdelete(clt *core.SDKClient, req *wdk.AlibabapricepromotionitemdeleteAPIRequest, session string) (*wdk.AlibabapricepromotionitemdeleteAPIResponse, error) {
	var resp wdk.AlibabapricepromotionitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
