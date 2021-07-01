package omniorder

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/omniorder"
)

/* 
库存初始化接口 
taobao.jst.astrolabe.storeinventory.iteminitial

ERP调用奇门的接口，对门店的库存进行初始化
*/
func TaobaoJstAstrolabeStoreinventoryIteminitial(clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest, session string) (*omniorder.TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse, error) {
    var resp omniorder.TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
