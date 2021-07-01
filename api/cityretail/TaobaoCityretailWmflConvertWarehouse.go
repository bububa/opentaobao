package cityretail

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cityretail"
)

/* 
同城零售完美履约转仓 
taobao.cityretail.wmfl.convert.warehouse

同城零售完美履约转仓
*/
func TaobaoCityretailWmflConvertWarehouse(clt *core.SDKClient, req *cityretail.TaobaoCityretailWmflConvertWarehouseAPIRequest, session string) (*cityretail.TaobaoCityretailWmflConvertWarehouseAPIResponse, error) {
    var resp cityretail.TaobaoCityretailWmflConvertWarehouseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
