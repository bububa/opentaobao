package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoFenxiaoDealerRequisitionorderGet
批量查询采购申请/经销采购单
taobao.fenxiao.dealer.requisitionorder.get

批量查询采购申请/经销采购单，目前支持供应商和分销商查询 */
func TaobaoFenxiaoDealerRequisitionorderGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoDealerRequisitionorderGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoDealerRequisitionorderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
