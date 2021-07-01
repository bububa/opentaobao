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
type TaobaoFenxiaoProductcatDeleteAPIRequest struct {
    model.Params
    // 产品线ID
    _productLineId   int64
}

// 初始化TaobaoFenxiaoProductcatDeleteAPIRequest对象
func NewTaobaoFenxiaoProductcatDeleteRequest() *TaobaoFenxiaoProductcatDeleteAPIRequest{
    return &TaobaoFenxiaoProductcatDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.productcat.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductLineId Setter
// 产品线ID
func (r *TaobaoFenxiaoProductcatDeleteAPIRequest) SetProductLineId(_productLineId int64) error {
    r._productLineId = _productLineId
    r.Set("product_line_id", _productLineId)
    return nil
}

// ProductLineId Getter
func (r TaobaoFenxiaoProductcatDeleteAPIRequest) GetProductLineId() int64 {
    return r._productLineId
}
