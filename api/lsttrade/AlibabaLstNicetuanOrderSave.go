package lsttrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lsttrade"
)

/* 
十荟团订单同步至零售通 
alibaba.lst.nicetuan.order.save

十荟团订单同步至零售通，十荟团单向写到零售通
*/
func AlibabaLstNicetuanOrderSave(clt *core.SDKClient, req *lsttrade.AlibabaLstNicetuanOrderSaveRequest, session string) (*lsttrade.AlibabaLstNicetuanOrderSaveAPIResponse, error) {
    var resp lsttrade.AlibabaLstNicetuanOrderSaveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
