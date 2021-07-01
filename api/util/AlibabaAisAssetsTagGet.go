package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
基础设施资产标签获取 
alibaba.ais.assets.tag.get

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
*/
func AlibabaAisAssetsTagGet(clt *core.SDKClient, req *util.AlibabaAisAssetsTagGetAPIRequest, session string) (*util.AlibabaAisAssetsTagGetAPIResponse, error) {
    var resp util.AlibabaAisAssetsTagGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
