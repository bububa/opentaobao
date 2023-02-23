package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItempropsGetAPIRequest 获取标准商品类目属性 API请求
// taobao.itemprops.get
//
// 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。
type TaobaoItempropsGetAPIRequest struct {
	model.Params
	// 属性的Key，支持多条，以“,”分隔
	_attrKeys []string
	// 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
	_fields string
	// 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
	_childPath string
	// 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
	_cid int64
	// 属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
	_pid int64
	// 父属性ID
	_parentPid int64
	// 获取类目的类型：1代表集市、2代表天猫
	_type int64
	// 是否关键属性。可选值:true(是),false(否)
	_isKeyProp bool
	// 是否销售属性。可选值:true(是),false(否)
	_isSaleProp bool
	// 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
	_isColorProp bool
	// 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
	_isEnumProp bool
	// 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
	_isInputProp bool
	// 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
	_isItemProp bool
}

// NewTaobaoItempropsGetRequest 初始化TaobaoItempropsGetAPIRequest对象
func NewTaobaoItempropsGetRequest() *TaobaoItempropsGetAPIRequest {
	return &TaobaoItempropsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItempropsGetAPIRequest) GetApiMethodName() string {
	return "taobao.itemprops.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItempropsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItempropsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttrKeys is AttrKeys Setter
// 属性的Key，支持多条，以“,”分隔
func (r *TaobaoItempropsGetAPIRequest) SetAttrKeys(_attrKeys []string) error {
	r._attrKeys = _attrKeys
	r.Set("attr_keys", _attrKeys)
	return nil
}

// GetAttrKeys AttrKeys Getter
func (r TaobaoItempropsGetAPIRequest) GetAttrKeys() []string {
	return r._attrKeys
}

// SetFields is Fields Setter
// 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
func (r *TaobaoItempropsGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoItempropsGetAPIRequest) GetFields() string {
	return r._fields
}

// SetChildPath is ChildPath Setter
// 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
func (r *TaobaoItempropsGetAPIRequest) SetChildPath(_childPath string) error {
	r._childPath = _childPath
	r.Set("child_path", _childPath)
	return nil
}

// GetChildPath ChildPath Getter
func (r TaobaoItempropsGetAPIRequest) GetChildPath() string {
	return r._childPath
}

// SetCid is Cid Setter
// 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
func (r *TaobaoItempropsGetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoItempropsGetAPIRequest) GetCid() int64 {
	return r._cid
}

// SetPid is Pid Setter
// 属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
func (r *TaobaoItempropsGetAPIRequest) SetPid(_pid int64) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaoItempropsGetAPIRequest) GetPid() int64 {
	return r._pid
}

// SetParentPid is ParentPid Setter
// 父属性ID
func (r *TaobaoItempropsGetAPIRequest) SetParentPid(_parentPid int64) error {
	r._parentPid = _parentPid
	r.Set("parent_pid", _parentPid)
	return nil
}

// GetParentPid ParentPid Getter
func (r TaobaoItempropsGetAPIRequest) GetParentPid() int64 {
	return r._parentPid
}

// SetType is Type Setter
// 获取类目的类型：1代表集市、2代表天猫
func (r *TaobaoItempropsGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoItempropsGetAPIRequest) GetType() int64 {
	return r._type
}

// SetIsKeyProp is IsKeyProp Setter
// 是否关键属性。可选值:true(是),false(否)
func (r *TaobaoItempropsGetAPIRequest) SetIsKeyProp(_isKeyProp bool) error {
	r._isKeyProp = _isKeyProp
	r.Set("is_key_prop", _isKeyProp)
	return nil
}

// GetIsKeyProp IsKeyProp Getter
func (r TaobaoItempropsGetAPIRequest) GetIsKeyProp() bool {
	return r._isKeyProp
}

// SetIsSaleProp is IsSaleProp Setter
// 是否销售属性。可选值:true(是),false(否)
func (r *TaobaoItempropsGetAPIRequest) SetIsSaleProp(_isSaleProp bool) error {
	r._isSaleProp = _isSaleProp
	r.Set("is_sale_prop", _isSaleProp)
	return nil
}

// GetIsSaleProp IsSaleProp Getter
func (r TaobaoItempropsGetAPIRequest) GetIsSaleProp() bool {
	return r._isSaleProp
}

// SetIsColorProp is IsColorProp Setter
// 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
func (r *TaobaoItempropsGetAPIRequest) SetIsColorProp(_isColorProp bool) error {
	r._isColorProp = _isColorProp
	r.Set("is_color_prop", _isColorProp)
	return nil
}

// GetIsColorProp IsColorProp Getter
func (r TaobaoItempropsGetAPIRequest) GetIsColorProp() bool {
	return r._isColorProp
}

// SetIsEnumProp is IsEnumProp Setter
// 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
func (r *TaobaoItempropsGetAPIRequest) SetIsEnumProp(_isEnumProp bool) error {
	r._isEnumProp = _isEnumProp
	r.Set("is_enum_prop", _isEnumProp)
	return nil
}

// GetIsEnumProp IsEnumProp Getter
func (r TaobaoItempropsGetAPIRequest) GetIsEnumProp() bool {
	return r._isEnumProp
}

// SetIsInputProp is IsInputProp Setter
// 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
func (r *TaobaoItempropsGetAPIRequest) SetIsInputProp(_isInputProp bool) error {
	r._isInputProp = _isInputProp
	r.Set("is_input_prop", _isInputProp)
	return nil
}

// GetIsInputProp IsInputProp Getter
func (r TaobaoItempropsGetAPIRequest) GetIsInputProp() bool {
	return r._isInputProp
}

// SetIsItemProp is IsItemProp Setter
// 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
func (r *TaobaoItempropsGetAPIRequest) SetIsItemProp(_isItemProp bool) error {
	r._isItemProp = _isItemProp
	r.Set("is_item_prop", _isItemProp)
	return nil
}

// GetIsItemProp IsItemProp Getter
func (r TaobaoItempropsGetAPIRequest) GetIsItemProp() bool {
	return r._isItemProp
}
