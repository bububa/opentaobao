package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodealerrequisitionorderremarkupdate 修改经销采购单备注
// taobao.fenxiao.dealer.requisitionorder.remark.update
//
// 供应商修改经销采购单备注
func Taobaofenxiaodealerrequisitionorderremarkupdate(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest, session string) (*fenxiao.TaobaofenxiaodealerrequisitionorderremarkupdateAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodealerrequisitionorderremarkupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
