package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItempropvaluesGetAPIRequest 获取标准类目属性值 API请求
// taobao.itempropvalues.get
//
// 获取标准类目属性值
type TaobaoItempropvaluesGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
	_fields []string
	// 属性的Key，支持多条，以“,”分隔
	_attrKeys []string
	// 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
	_pvs string
	// 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
	_cid int64
	// 获取类目的类型：1代表集市、2代表天猫
	_type int64
}

// NewTaobaoItempropvaluesGetRequest 初始化TaobaoItempropvaluesGetAPIRequest对象
func NewTaobaoItempropvaluesGetRequest() *TaobaoItempropvaluesGetAPIRequest {
	return &TaobaoItempropvaluesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItempropvaluesGetAPIRequest) GetApiMethodName() string {
	return "taobao.itempropvalues.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItempropvaluesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order
func (r *TaobaoItempropvaluesGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoItempropvaluesGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetAttrKeys is AttrKeys Setter
// 属性的Key，支持多条，以“,”分隔
func (r *TaobaoItempropvaluesGetAPIRequest) SetAttrKeys(_attrKeys []string) error {
	r._attrKeys = _attrKeys
	r.Set("attr_keys", _attrKeys)
	return nil
}

// GetAttrKeys AttrKeys Getter
func (r TaobaoItempropvaluesGetAPIRequest) GetAttrKeys() []string {
	return r._attrKeys
}

// SetPvs is Pvs Setter
// 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)
func (r *TaobaoItempropvaluesGetAPIRequest) SetPvs(_pvs string) error {
	r._pvs = _pvs
	r.Set("pvs", _pvs)
	return nil
}

// GetPvs Pvs Getter
func (r TaobaoItempropvaluesGetAPIRequest) GetPvs() string {
	return r._pvs
}

// SetCid is Cid Setter
// 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID
func (r *TaobaoItempropvaluesGetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoItempropvaluesGetAPIRequest) GetCid() int64 {
	return r._cid
}

// SetType is Type Setter
// 获取类目的类型：1代表集市、2代表天猫
func (r *TaobaoItempropvaluesGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoItempropvaluesGetAPIRequest) GetType() int64 {
	return r._type
}
