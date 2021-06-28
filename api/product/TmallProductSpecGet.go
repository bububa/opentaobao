package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
根据产品规格的Id号获取当个的规格信息 
tmall.product.spec.get

通过当个的spec_id获取到这个产品规格的信息，主要是因为产品规格是要经过审核的，所以通过这个接口可以获取到是否通过审核<br/>通过参看这个ProductSpec的status判断：<br/>1:表示审核通过<br/>3:表示等待审核。<br/>如果你的id找不到数据，那么就是审核被拒绝。
*/
func TmallProductSpecGet(clt *core.SDKClient, req *product.TmallProductSpecGetRequest, session string) (*product.TmallProductSpecGetAPIResponse, error) {
    var resp product.TmallProductSpecGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
