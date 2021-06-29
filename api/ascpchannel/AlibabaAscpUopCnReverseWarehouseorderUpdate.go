package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
供应链中台逆向入库单修改服务 
alibaba.ascp.uop.cn.reverse.warehouseorder.update

供应链中台逆向入库单修改服务
*/
func AlibabaAscpUopCnReverseWarehouseorderUpdate(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopCnReverseWarehouseorderUpdateRequest, session string) (*ascpchannel.AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
