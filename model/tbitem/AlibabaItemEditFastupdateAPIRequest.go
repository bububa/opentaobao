package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemEditFastupdateAPIRequest 商品编辑增量更新 API请求
// alibaba.item.edit.fastupdate
//
// 商品编辑增量更新;
// &lt;br/&gt;该接口编辑sku，只能更新价格、库存等信息，不能新增sku;
// &lt;br/&gt;新增sku用全量接口alibaba.item.edit.submit，先设置销售属性;
type AlibabaItemEditFastupdateAPIRequest struct {
	model.Params
	// 编辑后的schema信息(增量更新，只填写需要更新的字段)
	_schema string
	// 商品ID
	_itemId int64
	// 商品类目ID。若不需要修改商品类目，则不用填写
	_catId int64
	// 产品ID，若不需要修改关联的产品信息，则不需要填写
	_spuId int64
}

// NewAlibabaItemEditFastupdateRequest 初始化AlibabaItemEditFastupdateAPIRequest对象
func NewAlibabaItemEditFastupdateRequest() *AlibabaItemEditFastupdateAPIRequest {
	return &AlibabaItemEditFastupdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItemEditFastupdateAPIRequest) Reset() {
	r._schema = ""
	r._itemId = 0
	r._catId = 0
	r._spuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemEditFastupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.item.edit.fastupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemEditFastupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItemEditFastupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 编辑后的schema信息(增量更新，只填写需要更新的字段)
func (r *AlibabaItemEditFastupdateAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetSchema() string {
	return r._schema
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaItemEditFastupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCatId is CatId Setter
// 商品类目ID。若不需要修改商品类目，则不用填写
func (r *AlibabaItemEditFastupdateAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetSpuId is SpuId Setter
// 产品ID，若不需要修改关联的产品信息，则不需要填写
func (r *AlibabaItemEditFastupdateAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r AlibabaItemEditFastupdateAPIRequest) GetSpuId() int64 {
	return r._spuId
}

var poolAlibabaItemEditFastupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItemEditFastupdateRequest()
	},
}

// GetAlibabaItemEditFastupdateRequest 从 sync.Pool 获取 AlibabaItemEditFastupdateAPIRequest
func GetAlibabaItemEditFastupdateAPIRequest() *AlibabaItemEditFastupdateAPIRequest {
	return poolAlibabaItemEditFastupdateAPIRequest.Get().(*AlibabaItemEditFastupdateAPIRequest)
}

// ReleaseAlibabaItemEditFastupdateAPIRequest 将 AlibabaItemEditFastupdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaItemEditFastupdateAPIRequest(v *AlibabaItemEditFastupdateAPIRequest) {
	v.Reset()
	poolAlibabaItemEditFastupdateAPIRequest.Put(v)
}
