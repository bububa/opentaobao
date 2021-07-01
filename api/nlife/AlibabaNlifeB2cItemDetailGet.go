package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
b2c码详情查询 
alibaba.nlife.b2c.item.detail.get

根据零售+平台生成的唯一码获取对应详情
*/
func AlibabaNlifeB2cItemDetailGet(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cItemDetailGetAPIRequest, session string) (*nlife.AlibabaNlifeB2cItemDetailGetAPIResponse, error) {
    var resp nlife.AlibabaNlifeB2cItemDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
