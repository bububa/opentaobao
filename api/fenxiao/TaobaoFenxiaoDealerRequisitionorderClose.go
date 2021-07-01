package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
供应商/分销商关闭采购申请/经销采购单 
taobao.fenxiao.dealer.requisitionorder.close

供应商或分销商关闭采购申请/经销采购单
*/
func TaobaoFenxiaoDealerRequisitionorderClose(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest, session string) (*fenxiao.TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
