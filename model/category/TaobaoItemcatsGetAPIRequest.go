package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemcatsGetAPIRequest
获取后台供卖家发布商品的标准商品类目 API请求
taobao.itemcats.get

获取后台供卖家发布商品的标准商品类目。 */
type TaobaoItemcatsGetAPIRequest struct {
	model.Params
	// 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
	_cids []int64
	// 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
	_fields []string
	// 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
	_parentCid int64
}

// NewTaobaoItemcatsGetRequest 初始化TaobaoItemcatsGetAPIRequest对象
func NewTaobaoItemcatsGetRequest() *TaobaoItemcatsGetAPIRequest {
	return &TaobaoItemcatsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemcatsGetAPIRequest) GetApiMethodName() string {
	return "taobao.itemcats.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemcatsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Cids Setter
// 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
func (r *TaobaoItemcatsGetAPIRequest) SetCids(_cids []int64) error {
	r._cids = _cids
	r.Set("cids", _cids)
	return nil
}

// Get Cids Getter
func (r TaobaoItemcatsGetAPIRequest) GetCids() []int64 {
	return r._cids
}

// Set is Fields Setter
// 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
func (r *TaobaoItemcatsGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r TaobaoItemcatsGetAPIRequest) GetFields() []string {
	return r._fields
}

// Set is ParentCid Setter
// 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
func (r *TaobaoItemcatsGetAPIRequest) SetParentCid(_parentCid int64) error {
	r._parentCid = _parentCid
	r.Set("parent_cid", _parentCid)
	return nil
}

// Get ParentCid Getter
func (r TaobaoItemcatsGetAPIRequest) GetParentCid() int64 {
	return r._parentCid
}
