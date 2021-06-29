package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品下架 API请求
taobao.item.update.delisting

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateDelistingRequest struct {
    model.Params
    // 商品数字ID，该参数必须
    _numIid   int64
}

// 初始化TaobaoItemUpdateDelistingRequest对象
func NewTaobaoItemUpdateDelistingRequest() *TaobaoItemUpdateDelistingRequest{
    return &TaobaoItemUpdateDelistingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateDelistingRequest) GetApiMethodName() string {
    return "taobao.item.update.delisting"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateDelistingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemUpdateDelistingRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemUpdateDelistingRequest) GetNumIid() int64 {
    return r._numIid
}
