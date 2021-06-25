package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除产品线 APIRequest
taobao.fenxiao.productcat.delete

删除产品线
*/
type TaobaoFenxiaoProductcatDeleteRequest struct {
    model.Params

    // 产品线ID
    productLineId   int64 

}

func NewTaobaoFenxiaoProductcatDeleteRequest() *TaobaoFenxiaoProductcatDeleteRequest{
    return &TaobaoFenxiaoProductcatDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductcatDeleteRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcat.delete"
}

func (r TaobaoFenxiaoProductcatDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductcatDeleteRequest) SetProductLineId(productLineId int64) error {
    r.productLineId = productLineId
    r.Set("product_line_id", productLineId)
    return nil
}

func (r TaobaoFenxiaoProductcatDeleteRequest) GetProductLineId() int64 {
    return r.productLineId
}

