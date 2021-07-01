package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoFenxiaoDealerRequisitionorderAgree
供应商/分销商通过采购申请/经销采购单申请
taobao.fenxiao.dealer.requisitionorder.agree

供应商或分销商通过采购申请/经销采购单审核 */
func TaobaoFenxiaoDealerRequisitionorderAgree(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest, session string) (*fenxiao.TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
