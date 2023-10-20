package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodistributorproductquantityget 分销商查询产品库存
// taobao.fenxiao.distributor.product.quantity.get
//
// 分销商查询产品库存
func Taobaofenxiaodistributorproductquantityget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodistributorproductquantitygetAPIRequest, session string) (*fenxiao.TaobaofenxiaodistributorproductquantitygetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodistributorproductquantitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
