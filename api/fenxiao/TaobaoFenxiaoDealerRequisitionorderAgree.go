package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaodealerrequisitionorderagree 供应商/分销商通过采购申请/经销采购单申请
// taobao.fenxiao.dealer.requisitionorder.agree
//
// 供应商或分销商通过采购申请/经销采购单审核
func Taobaofenxiaodealerrequisitionorderagree(clt *core.SDKClient, req *fenxiao.TaobaofenxiaodealerrequisitionorderagreeAPIRequest, session string) (*fenxiao.TaobaofenxiaodealerrequisitionorderagreeAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaodealerrequisitionorderagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
