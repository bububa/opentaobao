package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
产品图片删除 
taobao.fenxiao.product.image.delete

产品图片删除，只删除图片信息，不真正删除图片
*/
func TaobaoFenxiaoProductImageDelete(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductImageDeleteAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductImageDeleteAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductImageDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
