package txcs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/txcs"
)

/* 
天猫超市外部商家财务账单信息查询 
tmall.txcs.finance.bill.query

提供天猫超市外部合作商家财务账单对账
*/
func TmallTxcsFinanceBillQuery(clt *core.SDKClient, req *txcs.TmallTxcsFinanceBillQueryAPIRequest, session string) (*txcs.TmallTxcsFinanceBillQueryAPIResponse, error) {
    var resp txcs.TmallTxcsFinanceBillQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
