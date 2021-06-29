package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
taobao.item.update.delisting.tmall API请求
taobao.item.update.delisting.tmall

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateDelistingTmallRequest struct {
    model.Params
    // 商品数字ID，该参数必须
    _numIid   int64
}

// 初始化TaobaoItemUpdateDelistingTmallRequest对象
func NewTaobaoItemUpdateDelistingTmallRequest() *TaobaoItemUpdateDelistingTmallRequest{
    return &TaobaoItemUpdateDelistingTmallRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateDelistingTmallRequest) GetApiMethodName() string {
    return "taobao.item.update.delisting.tmall"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateDelistingTmallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemUpdateDelistingTmallRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemUpdateDelistingTmallRequest) GetNumIid() int64 {
    return r._numIid
}
