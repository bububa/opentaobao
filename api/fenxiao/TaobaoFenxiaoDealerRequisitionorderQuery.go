package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodealerrequisitionorderquery 按编号查询采购申请/经销采购单
// taobao.fenxiao.dealer.requisitionorder.query
//
// 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
func Taobaofenxiaodealerrequisitionorderquery(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodealerrequisitionorderqueryAPIRequest, session string) (*fenxiao.TaobaofenxiaodealerrequisitionorderqueryAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodealerrequisitionorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
