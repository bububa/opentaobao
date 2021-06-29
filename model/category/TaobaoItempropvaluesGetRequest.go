package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标准类目属性值 API请求
taobao.itempropvalues.get

获取标准类目属性值
*/
type TaobaoItempropvaluesGetRequest struct {
    model.Params
    // 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
    _fields   []string
    // 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
    _cid   int64
    // 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
    _pvs   string
    // 获取类目的类型：1代表集市、2代表天猫
    _type   int64
    // 属性的Key，支持多条，以“,”分隔
    _attrKeys   []string
}

// 初始化TaobaoItempropvaluesGetRequest对象
func NewTaobaoItempropvaluesGetRequest() *TaobaoItempropvaluesGetRequest{
    return &TaobaoItempropvaluesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItempropvaluesGetRequest) GetApiMethodName() string {
    return "taobao.itempropvalues.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItempropvaluesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
func (r *TaobaoItempropvaluesGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoItempropvaluesGetRequest) GetFields() []string {
    return r._fields
}
// Cid Setter
// 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
func (r *TaobaoItempropvaluesGetRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TaobaoItempropvaluesGetRequest) GetCid() int64 {
    return r._cid
}
// Pvs Setter
// 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
func (r *TaobaoItempropvaluesGetRequest) SetPvs(_pvs string) error {
    r._pvs = _pvs
    r.Set("pvs", _pvs)
    return nil
}

// Pvs Getter
func (r TaobaoItempropvaluesGetRequest) GetPvs() string {
    return r._pvs
}
// Type Setter
// 获取类目的类型：1代表集市、2代表天猫
func (r *TaobaoItempropvaluesGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoItempropvaluesGetRequest) GetType() int64 {
    return r._type
}
// AttrKeys Setter
// 属性的Key，支持多条，以“,”分隔
func (r *TaobaoItempropvaluesGetRequest) SetAttrKeys(_attrKeys []string) error {
    r._attrKeys = _attrKeys
    r.Set("attr_keys", _attrKeys)
    return nil
}

// AttrKeys Getter
func (r TaobaoItempropvaluesGetRequest) GetAttrKeys() []string {
    return r._attrKeys
}
