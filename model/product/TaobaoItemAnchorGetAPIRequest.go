package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取可用宝贝描述规范化模块 API请求
taobao.item.anchor.get

根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点
*/
type TaobaoItemAnchorGetAPIRequest struct {
    model.Params
    // 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.(最小值：-1，最大值：1)
    _type   int64
    // 对应类目编号
    _catId   int64
}

// 初始化TaobaoItemAnchorGetAPIRequest对象
func NewTaobaoItemAnchorGetRequest() *TaobaoItemAnchorGetAPIRequest{
    return &TaobaoItemAnchorGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemAnchorGetAPIRequest) GetApiMethodName() string {
    return "taobao.item.anchor.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemAnchorGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Type Setter
// 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.(最小值：-1，最大值：1)
func (r *TaobaoItemAnchorGetAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoItemAnchorGetAPIRequest) GetType() int64 {
    return r._type
}
// CatId Setter
// 对应类目编号
func (r *TaobaoItemAnchorGetAPIRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r TaobaoItemAnchorGetAPIRequest) GetCatId() int64 {
    return r._catId
}
