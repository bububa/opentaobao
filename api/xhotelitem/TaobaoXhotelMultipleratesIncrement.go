package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
复杂房价推送接口（批量增量） 
taobao.xhotel.multiplerates.increment

复杂房价批量增量更新，只会更新指定日期的信息
完全涵盖了taobao.xhotel.rates.increment接口的功能
*/
func TaobaoXhotelMultipleratesIncrement(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultipleratesIncrementRequest, session string) (*xhotelitem.TaobaoXhotelMultipleratesIncrementAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelMultipleratesIncrementAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
