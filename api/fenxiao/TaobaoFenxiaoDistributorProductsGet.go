package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodistributorproductsget 分销商查询产品信息
// taobao.fenxiao.distributor.products.get
//
// 分销商查询供应商产品信息
func Taobaofenxiaodistributorproductsget(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodistributorproductsgetAPIRequest, session string) (*fenxiao.TaobaofenxiaodistributorproductsgetAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodistributorproductsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
