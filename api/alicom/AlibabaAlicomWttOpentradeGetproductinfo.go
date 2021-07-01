package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
查询存送产品信息 
alibaba.alicom.wtt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
func AlibabaAlicomWttOpentradeGetproductinfo(clt *core.SDKClient, req *alicom.AlibabaAlicomWttOpentradeGetproductinfoAPIRequest, session string) (*alicom.AlibabaAlicomWttOpentradeGetproductinfoAPIResponse, error) {
    var resp alicom.AlibabaAlicomWttOpentradeGetproductinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
