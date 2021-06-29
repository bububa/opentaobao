package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除属性图片 API请求
taobao.item.propimg.delete

删除propimg_id 所指定的商品属性图片 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>propimg_id对应的属性图片需要属于num_iid对应的商品
*/
type TaobaoItemPropimgDeleteRequest struct {
    model.Params
    // 商品属性图片ID
    _id   int64
    // 商品数字ID，必选
    _numIid   int64
}

// 初始化TaobaoItemPropimgDeleteRequest对象
func NewTaobaoItemPropimgDeleteRequest() *TaobaoItemPropimgDeleteRequest{
    return &TaobaoItemPropimgDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemPropimgDeleteRequest) GetApiMethodName() string {
    return "taobao.item.propimg.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemPropimgDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 商品属性图片ID
func (r *TaobaoItemPropimgDeleteRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoItemPropimgDeleteRequest) GetId() int64 {
    return r._id
}
// NumIid Setter
// 商品数字ID，必选
func (r *TaobaoItemPropimgDeleteRequest) SetNumIid(_numIid int64) error {
    r._numIid = _numIid
    r.Set("num_iid", _numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemPropimgDeleteRequest) GetNumIid() int64 {
    return r._numIid
}
