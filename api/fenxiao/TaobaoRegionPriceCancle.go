package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
取消区域价格 
taobao.region.price.cancle

取消区域价格
*/
func TaobaoRegionPriceCancle(clt *core.SDKClient, req *fenxiao.TaobaoRegionPriceCancleAPIRequest, session string) (*fenxiao.TaobaoRegionPriceCancleAPIResponse, error) {
    var resp fenxiao.TaobaoRegionPriceCancleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
