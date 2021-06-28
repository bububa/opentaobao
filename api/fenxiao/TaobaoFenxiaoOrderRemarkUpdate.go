package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
修改采购单备注 
taobao.fenxiao.order.remark.update

供应商修改采购单备注
*/
func TaobaoFenxiaoOrderRemarkUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrderRemarkUpdateRequest, session string) (*fenxiao.TaobaoFenxiaoOrderRemarkUpdateAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoOrderRemarkUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
