package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdate 修改经销采购单备注
// taobao.fenxiao.dealer.requisitionorder.remark.update
//
// 供应商修改经销采购单备注
func TaobaoFenxiaoDealerRequisitionorderRemarkUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest, session string) (*fenxiao.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
