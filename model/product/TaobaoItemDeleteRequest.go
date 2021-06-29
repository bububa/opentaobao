package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单条商品 API请求
taobao.item.delete

删除单条商品
*/
type TaobaoItemDeleteRequest struct {
    model.Params
    // 商品数字ID，该参数必须
    _numIid   int64
}

// 初始化TaobaoItemDeleteRequest对象
func NewTaobaoItemDeleteRequest() *TaobaoItemDeleteRequest{
    return &TaobaoItemDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemDeleteRequest) GetApiMethodName() string {
    return "taobao.item.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemDeleteRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemDeleteRequest) GetNumIid() int64 {
    return r._numIid
}
