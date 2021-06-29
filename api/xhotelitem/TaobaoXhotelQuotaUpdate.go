package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
库存更新接口 
taobao.xhotel.quota.update

库存更新接口
*/
func TaobaoXhotelQuotaUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelQuotaUpdateRequest, session string) (*xhotelitem.TaobaoXhotelQuotaUpdateAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelQuotaUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
