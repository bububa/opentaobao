package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除产品线 API请求
taobao.fenxiao.productcat.delete

删除产品线
*/
type TaobaoFenxiaoProductcatDeleteRequest struct {
    model.Params
    // 产品线ID
    productLineId   int64
}

// 初始化TaobaoFenxiaoProductcatDeleteRequest对象
func NewTaobaoFenxiaoProductcatDeleteRequest() *TaobaoFenxiaoProductcatDeleteRequest{
    return &TaobaoFenxiaoProductcatDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatDeleteRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcat.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductLineId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductcatDeleteRequest) SetProductLineId(productLineId int64) error {
    r.productLineId = productLineId
    r.Set("product_line_id", productLineId)
    return nil
}

// ProductLineId Getter
func (r TaobaoFenxiaoProductcatDeleteRequest) GetProductLineId() int64 {
    return r.productLineId
}
