package category

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemcatsGetAPIRequest 获取后台供卖家发布商品的标准商品类目 API请求
// taobao.itemcats.get
//
// 获取后台供卖家发布商品的标准商品类目。
type TaobaoItemcatsGetAPIRequest struct {
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

// NewTaobaoItemcatsGetRequest 初始化TaobaoItemcatsGetAPIRequest对象
func NewTaobaoItemcatsGetRequest() *TaobaoItemcatsGetAPIRequest {
	return &TaobaoItemcatsGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemcatsGetAPIRequest) Reset() {
	r._cids = r._cids[:0]
	r._fields = r._fields[:0]
	r._datetime = ""
	r._parentCid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemcatsGetAPIRequest) GetApiMethodName() string {
	return "taobao.itemcats.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemcatsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemcatsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCids is Cids Setter
// 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
func (r *TaobaoItemcatsGetAPIRequest) SetCids(_cids []int64) error {
	r._cids = _cids
	r.Set("cids", _cids)
	return nil
}

// GetCids Cids Getter
func (r TaobaoItemcatsGetAPIRequest) GetCids() []int64 {
	return r._cids
}

// SetFields is Fields Setter
// 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
func (r *TaobaoItemcatsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoItemcatsGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetDatetime is Datetime Setter
// 无效字段，暂无法使用。时间戳（格式:yyyy-MM-dd HH:mm:ss）如果该字段没有传，则取当前所有的类目信息,如果传了parent_cid或者cids，则忽略datetime，如果该字段传了，那么可以查询到该时间到现在为止的增量变化
func (r *TaobaoItemcatsGetAPIRequest) SetDatetime(_datetime string) error {
	r._datetime = _datetime
	r.Set("datetime", _datetime)
	return nil
}

// GetDatetime Datetime Getter
func (r TaobaoItemcatsGetAPIRequest) GetDatetime() string {
	return r._datetime
}

// SetParentCid is ParentCid Setter
// 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
func (r *TaobaoItemcatsGetAPIRequest) SetParentCid(_parentCid int64) error {
	r._parentCid = _parentCid
	r.Set("parent_cid", _parentCid)
	return nil
}

// GetParentCid ParentCid Getter
func (r TaobaoItemcatsGetAPIRequest) GetParentCid() int64 {
	return r._parentCid
}

var poolTaobaoItemcatsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemcatsGetRequest()
	},
}

// GetTaobaoItemcatsGetRequest 从 sync.Pool 获取 TaobaoItemcatsGetAPIRequest
func GetTaobaoItemcatsGetAPIRequest() *TaobaoItemcatsGetAPIRequest {
	return poolTaobaoItemcatsGetAPIRequest.Get().(*TaobaoItemcatsGetAPIRequest)
}

// ReleaseTaobaoItemcatsGetAPIRequest 将 TaobaoItemcatsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemcatsGetAPIRequest(v *TaobaoItemcatsGetAPIRequest) {
	v.Reset()
	poolTaobaoItemcatsGetAPIRequest.Put(v)
}
