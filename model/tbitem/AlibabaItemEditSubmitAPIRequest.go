package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemEditSubmitAPIRequest 商品编辑提交schema信息 API请求
// alibaba.item.edit.submit
//
// 商品编辑提交schema信息
type AlibabaItemEditSubmitAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 编辑后的schema信息，通过alibaba.item.edit.schema.get获取
	_schema string
	// 商品ID
	_itemId int64
	// 商品类目ID。若不需要修改商品类目，则不用填写
	_catId int64
	// 产品ID，若不需要修改关联的产品信息，则不需要填写
	_spuId int64
}

// NewAlibabaItemEditSubmitRequest 初始化AlibabaItemEditSubmitAPIRequest对象
func NewAlibabaItemEditSubmitRequest() *AlibabaItemEditSubmitAPIRequest {
	return &AlibabaItemEditSubmitAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItemEditSubmitAPIRequest) Reset() {
	r._bizType = ""
	r._schema = ""
	r._itemId = 0
	r._catId = 0
	r._spuId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemEditSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.item.edit.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemEditSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItemEditSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizType is BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemEditSubmitAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaItemEditSubmitAPIRequest) GetBizType() string {
	return r._bizType
}

// SetSchema is Schema Setter
// 编辑后的schema信息，通过alibaba.item.edit.schema.get获取
func (r *AlibabaItemEditSubmitAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaItemEditSubmitAPIRequest) GetSchema() string {
	return r._schema
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaItemEditSubmitAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaItemEditSubmitAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCatId is CatId Setter
// 商品类目ID。若不需要修改商品类目，则不用填写
func (r *AlibabaItemEditSubmitAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaItemEditSubmitAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetSpuId is SpuId Setter
// 产品ID，若不需要修改关联的产品信息，则不需要填写
func (r *AlibabaItemEditSubmitAPIRequest) SetSpuId(_spuId int64) error {
	r._spuId = _spuId
	r.Set("spu_id", _spuId)
	return nil
}

// GetSpuId SpuId Getter
func (r AlibabaItemEditSubmitAPIRequest) GetSpuId() int64 {
	return r._spuId
}

var poolAlibabaItemEditSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItemEditSubmitRequest()
	},
}

// GetAlibabaItemEditSubmitRequest 从 sync.Pool 获取 AlibabaItemEditSubmitAPIRequest
func GetAlibabaItemEditSubmitAPIRequest() *AlibabaItemEditSubmitAPIRequest {
	return poolAlibabaItemEditSubmitAPIRequest.Get().(*AlibabaItemEditSubmitAPIRequest)
}

// ReleaseAlibabaItemEditSubmitAPIRequest 将 AlibabaItemEditSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaItemEditSubmitAPIRequest(v *AlibabaItemEditSubmitAPIRequest) {
	v.Reset()
	poolAlibabaItemEditSubmitAPIRequest.Put(v)
}
