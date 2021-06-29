package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
获取终端类型下品牌列表 
yunos.tvpubadmin.device.brands

获取终端类型下品牌列表
*/
func YunosTvpubadminDeviceBrands(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceBrandsRequest, session string) (*tvupadmin.YunosTvpubadminDeviceBrandsAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminDeviceBrandsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
