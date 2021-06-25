package dt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dt"
)

/* 
线索报价价格校验 
alibaba.dt.tmllcar.pricevalidate

根据选定的车型和城市，校验汽车价格是否通过
入参：车型ID，城市名称，价格
输出：N 校验失败，校验成功不返回值
*/
func AlibabaDtTmllcarPricevalidate(clt *core.SDKClient, req *dt.AlibabaDtTmllcarPricevalidateRequest, session string) (*dt.AlibabaDtTmllcarPricevalidateResponse, error) {
    var resp dt.AlibabaDtTmllcarPricevalidateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
