package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductgradepriceget 等级折扣查询
// taobao.fenxiao.product.gradeprice.get
//
// 等级折扣查询
func Taobaofenxiaoproductgradepriceget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductgradepricegetAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductgradepricegetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductgradepricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
