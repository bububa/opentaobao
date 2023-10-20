package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemcatsgetAPIRequest 获取后台供卖家发布商品的标准商品类目 API请求
// taobao.itemcats.get
//
// 获取后台供卖家发布商品的标准商品类目。
type TaobaoitemcatsgetAPIRequest struct {
	model.Params
	// 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
	_cids []int64
	// 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
	_fields []string
	// 无效字段，暂无法使用。时间戳（格式:yyyy-MM-dd HH:mm:ss）如果该字段没有传，则取当前所有的类目信息,如果传了parent_cid或者cids，则忽略datetime，如果该字段传了，那么可以查询到该时间到现在为止的增量变化
	_datetime string
	// 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
	_parentCid int64
}

// NewTaobaoitemcatsgetRequest 初始化TaobaoitemcatsgetAPIRequest对象
func NewTaobaoitemcatsgetRequest() *TaobaoitemcatsgetAPIRequest {
	return &TaobaoitemcatsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemcatsgetAPIRequest) GetApiMethodName() string {
	return "taobao.itemcats.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemcatsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemcatsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCids is Cids Setter
// 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
func (r *TaobaoitemcatsgetAPIRequest) SetCids(_cids []int64) error {
	r._cids = _cids
	r.Set("cids", _cids)
	return nil
}

// GetCids Cids Getter
func (r TaobaoitemcatsgetAPIRequest) GetCids() []int64 {
	return r._cids
}

// SetFields is Fields Setter
// 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
func (r *TaobaoitemcatsgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoitemcatsgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetDatetime is Datetime Setter
// 无效字段，暂无法使用。时间戳（格式:yyyy-MM-dd HH:mm:ss）如果该字段没有传，则取当前所有的类目信息,如果传了parent_cid或者cids，则忽略datetime，如果该字段传了，那么可以查询到该时间到现在为止的增量变化
func (r *TaobaoitemcatsgetAPIRequest) SetDatetime(_datetime string) error {
	r._datetime = _datetime
	r.Set("datetime", _datetime)
	return nil
}

// GetDatetime Datetime Getter
func (r TaobaoitemcatsgetAPIRequest) GetDatetime() string {
	return r._datetime
}

// SetParentCid is ParentCid Setter
// 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
func (r *TaobaoitemcatsgetAPIRequest) SetParentCid(_parentCid int64) error {
	r._parentCid = _parentCid
	r.Set("parent_cid", _parentCid)
	return nil
}

// GetParentCid ParentCid Getter
func (r TaobaoitemcatsgetAPIRequest) GetParentCid() int64 {
	return r._parentCid
}
