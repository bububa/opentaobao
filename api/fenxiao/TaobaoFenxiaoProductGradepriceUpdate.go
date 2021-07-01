package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
根据sku设置折扣价 
taobao.fenxiao.product.gradeprice.update

供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
*/
func TaobaoFenxiaoProductGradepriceUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductGradepriceUpdateAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductGradepriceUpdateAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductGradepriceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
