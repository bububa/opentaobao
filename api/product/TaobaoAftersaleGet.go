package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
查询用户售后服务模板 
taobao.aftersale.get

查询用户设置的售后服务模板，仅返回标题和id
*/
func TaobaoAftersaleGet(clt *core.SDKClient, req *product.TaobaoAftersaleGetAPIRequest, session string) (*product.TaobaoAftersaleGetAPIResponse, error) {
    var resp product.TaobaoAftersaleGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
