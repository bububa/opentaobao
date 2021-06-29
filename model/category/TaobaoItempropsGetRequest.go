package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标准商品类目属性 API请求
taobao.itemprops.get

通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
*/
type TaobaoItempropsGetRequest struct {
    model.Params
    // 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
    _fields   string
    // 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
    _cid   int64
    // 属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
    _pid   int64
    // 父属性ID
    _parentPid   int64
    // 是否关键属性。可选值:true(是),false(否)
    _isKeyProp   bool
    // 是否销售属性。可选值:true(是),false(否)
    _isSaleProp   bool
    // 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
    _isColorProp   bool
    // 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
    _isEnumProp   bool
    // 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
    _isInputProp   bool
    // 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
    _isItemProp   bool
    // 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
    _childPath   string
    // 获取类目的类型：1代表集市、2代表天猫
    _type   int64
    // 属性的Key，支持多条，以“,”分隔
    _attrKeys   []string
}

// 初始化TaobaoItempropsGetRequest对象
func NewTaobaoItempropsGetRequest() *TaobaoItempropsGetRequest{
    return &TaobaoItempropsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItempropsGetRequest) GetApiMethodName() string {
    return "taobao.itemprops.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItempropsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
func (r *TaobaoItempropsGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoItempropsGetRequest) GetFields() string {
    return r._fields
}
// Cid Setter
// 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
func (r *TaobaoItempropsGetRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TaobaoItempropsGetRequest) GetCid() int64 {
    return r._cid
}
// Pid Setter
// 属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
func (r *TaobaoItempropsGetRequest) SetPid(_pid int64) error {
    r._pid = _pid
    r.Set("pid", _pid)
    return nil
}

// Pid Getter
func (r TaobaoItempropsGetRequest) GetPid() int64 {
    return r._pid
}
// ParentPid Setter
// 父属性ID
func (r *TaobaoItempropsGetRequest) SetParentPid(_parentPid int64) error {
    r._parentPid = _parentPid
    r.Set("parent_pid", _parentPid)
    return nil
}

// ParentPid Getter
func (r TaobaoItempropsGetRequest) GetParentPid() int64 {
    return r._parentPid
}
// IsKeyProp Setter
// 是否关键属性。可选值:true(是),false(否)
func (r *TaobaoItempropsGetRequest) SetIsKeyProp(_isKeyProp bool) error {
    r._isKeyProp = _isKeyProp
    r.Set("is_key_prop", _isKeyProp)
    return nil
}

// IsKeyProp Getter
func (r TaobaoItempropsGetRequest) GetIsKeyProp() bool {
    return r._isKeyProp
}
// IsSaleProp Setter
// 是否销售属性。可选值:true(是),false(否)
func (r *TaobaoItempropsGetRequest) SetIsSaleProp(_isSaleProp bool) error {
    r._isSaleProp = _isSaleProp
    r.Set("is_sale_prop", _isSaleProp)
    return nil
}

// IsSaleProp Getter
func (r TaobaoItempropsGetRequest) GetIsSaleProp() bool {
    return r._isSaleProp
}
// IsColorProp Setter
// 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
func (r *TaobaoItempropsGetRequest) SetIsColorProp(_isColorProp bool) error {
    r._isColorProp = _isColorProp
    r.Set("is_color_prop", _isColorProp)
    return nil
}

// IsColorProp Getter
func (r TaobaoItempropsGetRequest) GetIsColorProp() bool {
    return r._isColorProp
}
// IsEnumProp Setter
// 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
func (r *TaobaoItempropsGetRequest) SetIsEnumProp(_isEnumProp bool) error {
    r._isEnumProp = _isEnumProp
    r.Set("is_enum_prop", _isEnumProp)
    return nil
}

// IsEnumProp Getter
func (r TaobaoItempropsGetRequest) GetIsEnumProp() bool {
    return r._isEnumProp
}
// IsInputProp Setter
// 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
func (r *TaobaoItempropsGetRequest) SetIsInputProp(_isInputProp bool) error {
    r._isInputProp = _isInputProp
    r.Set("is_input_prop", _isInputProp)
    return nil
}

// IsInputProp Getter
func (r TaobaoItempropsGetRequest) GetIsInputProp() bool {
    return r._isInputProp
}
// IsItemProp Setter
// 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
func (r *TaobaoItempropsGetRequest) SetIsItemProp(_isItemProp bool) error {
    r._isItemProp = _isItemProp
    r.Set("is_item_prop", _isItemProp)
    return nil
}

// IsItemProp Getter
func (r TaobaoItempropsGetRequest) GetIsItemProp() bool {
    return r._isItemProp
}
// ChildPath Setter
// 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
func (r *TaobaoItempropsGetRequest) SetChildPath(_childPath string) error {
    r._childPath = _childPath
    r.Set("child_path", _childPath)
    return nil
}

// ChildPath Getter
func (r TaobaoItempropsGetRequest) GetChildPath() string {
    return r._childPath
}
// Type Setter
// 获取类目的类型：1代表集市、2代表天猫
func (r *TaobaoItempropsGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoItempropsGetRequest) GetType() int64 {
    return r._type
}
// AttrKeys Setter
// 属性的Key，支持多条，以“,”分隔
func (r *TaobaoItempropsGetRequest) SetAttrKeys(_attrKeys []string) error {
    r._attrKeys = _attrKeys
    r.Set("attr_keys", _attrKeys)
    return nil
}

// AttrKeys Getter
func (r TaobaoItempropsGetRequest) GetAttrKeys() []string {
    return r._attrKeys
}
