package paimai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiItempropsGetAPIRequest 拍卖相关类目属性 API请求
// taobao.paimai.itemprops.get
//
// 读取拍卖相关类目属性
type TaobaoPaimaiItempropsGetAPIRequest struct {
	model.Params
	// 属性的Key，支持多条，以“,”分隔
	_attrKeys []string
	// 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
	_fields []string
	// 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
	_childPath string
	// 增量时间戳。格式:yyyy-MM-dd HH:mm:ss假如传2005-01-01 00:00:00，则取所有的属性和子属性ID(如果传了pid会忽略datetime)
	_datetime string
	// 获取类目的类型：1代表集市、2代表天猫
	_type int64
	// 属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
	_pid int64
	// 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
	_cid int64
	// 父属性ID
	_parentPid int64
	// 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
	_isItemProp bool
	// 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
	_isInputProp bool
	// 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
	_isEnumProp bool
	// 是否销售属性。可选值:true(是),false(否)
	_isSaleProp bool
	// 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
	_isColorProp bool
	// 是否关键属性。可选值:true(是),false(否)
	_isKeyProp bool
}

// NewTaobaoPaimaiItempropsGetRequest 初始化TaobaoPaimaiItempropsGetAPIRequest对象
func NewTaobaoPaimaiItempropsGetRequest() *TaobaoPaimaiItempropsGetAPIRequest {
	return &TaobaoPaimaiItempropsGetAPIRequest{
		Params: model.NewParams(14),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPaimaiItempropsGetAPIRequest) Reset() {
	r._attrKeys = r._attrKeys[:0]
	r._fields = r._fields[:0]
	r._childPath = ""
	r._datetime = ""
	r._type = 0
	r._pid = 0
	r._cid = 0
	r._parentPid = 0
	r._isItemProp = false
	r._isInputProp = false
	r._isEnumProp = false
	r._isSaleProp = false
	r._isColorProp = false
	r._isKeyProp = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPaimaiItempropsGetAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.itemprops.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPaimaiItempropsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPaimaiItempropsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttrKeys is AttrKeys Setter
// 属性的Key，支持多条，以“,”分隔
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetAttrKeys(_attrKeys []string) error {
	r._attrKeys = _attrKeys
	r.Set("attr_keys", _attrKeys)
	return nil
}

// GetAttrKeys AttrKeys Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetAttrKeys() []string {
	return r._attrKeys
}

// SetFields is Fields Setter
// 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetChildPath is ChildPath Setter
// 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetChildPath(_childPath string) error {
	r._childPath = _childPath
	r.Set("child_path", _childPath)
	return nil
}

// GetChildPath ChildPath Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetChildPath() string {
	return r._childPath
}

// SetDatetime is Datetime Setter
// 增量时间戳。格式:yyyy-MM-dd HH:mm:ss假如传2005-01-01 00:00:00，则取所有的属性和子属性ID(如果传了pid会忽略datetime)
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetDatetime(_datetime string) error {
	r._datetime = _datetime
	r.Set("datetime", _datetime)
	return nil
}

// GetDatetime Datetime Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetDatetime() string {
	return r._datetime
}

// SetType is Type Setter
// 获取类目的类型：1代表集市、2代表天猫
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetType() int64 {
	return r._type
}

// SetPid is Pid Setter
// 属性id (取类目属性时，传pid，不用同时传PID和parent_pid)
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetPid(_pid int64) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetPid() int64 {
	return r._pid
}

// SetCid is Cid Setter
// 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetCid() int64 {
	return r._cid
}

// SetParentPid is ParentPid Setter
// 父属性ID
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetParentPid(_parentPid int64) error {
	r._parentPid = _parentPid
	r.Set("parent_pid", _parentPid)
	return nil
}

// GetParentPid ParentPid Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetParentPid() int64 {
	return r._parentPid
}

// SetIsItemProp is IsItemProp Setter
// 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetIsItemProp(_isItemProp bool) error {
	r._isItemProp = _isItemProp
	r.Set("is_item_prop", _isItemProp)
	return nil
}

// GetIsItemProp IsItemProp Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetIsItemProp() bool {
	return r._isItemProp
}

// SetIsInputProp is IsInputProp Setter
// 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetIsInputProp(_isInputProp bool) error {
	r._isInputProp = _isInputProp
	r.Set("is_input_prop", _isInputProp)
	return nil
}

// GetIsInputProp IsInputProp Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetIsInputProp() bool {
	return r._isInputProp
}

// SetIsEnumProp is IsEnumProp Setter
// 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetIsEnumProp(_isEnumProp bool) error {
	r._isEnumProp = _isEnumProp
	r.Set("is_enum_prop", _isEnumProp)
	return nil
}

// GetIsEnumProp IsEnumProp Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetIsEnumProp() bool {
	return r._isEnumProp
}

// SetIsSaleProp is IsSaleProp Setter
// 是否销售属性。可选值:true(是),false(否)
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetIsSaleProp(_isSaleProp bool) error {
	r._isSaleProp = _isSaleProp
	r.Set("is_sale_prop", _isSaleProp)
	return nil
}

// GetIsSaleProp IsSaleProp Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetIsSaleProp() bool {
	return r._isSaleProp
}

// SetIsColorProp is IsColorProp Setter
// 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetIsColorProp(_isColorProp bool) error {
	r._isColorProp = _isColorProp
	r.Set("is_color_prop", _isColorProp)
	return nil
}

// GetIsColorProp IsColorProp Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetIsColorProp() bool {
	return r._isColorProp
}

// SetIsKeyProp is IsKeyProp Setter
// 是否关键属性。可选值:true(是),false(否)
func (r *TaobaoPaimaiItempropsGetAPIRequest) SetIsKeyProp(_isKeyProp bool) error {
	r._isKeyProp = _isKeyProp
	r.Set("is_key_prop", _isKeyProp)
	return nil
}

// GetIsKeyProp IsKeyProp Getter
func (r TaobaoPaimaiItempropsGetAPIRequest) GetIsKeyProp() bool {
	return r._isKeyProp
}

var poolTaobaoPaimaiItempropsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPaimaiItempropsGetRequest()
	},
}

// GetTaobaoPaimaiItempropsGetRequest 从 sync.Pool 获取 TaobaoPaimaiItempropsGetAPIRequest
func GetTaobaoPaimaiItempropsGetAPIRequest() *TaobaoPaimaiItempropsGetAPIRequest {
	return poolTaobaoPaimaiItempropsGetAPIRequest.Get().(*TaobaoPaimaiItempropsGetAPIRequest)
}

// ReleaseTaobaoPaimaiItempropsGetAPIRequest 将 TaobaoPaimaiItempropsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPaimaiItempropsGetAPIRequest(v *TaobaoPaimaiItempropsGetAPIRequest) {
	v.Reset()
	poolTaobaoPaimaiItempropsGetAPIRequest.Put(v)
}
