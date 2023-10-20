package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemschemaincrementupdateAPIRequest 天猫根据规则增量更新商品 API请求
// tmall.item.schema.increment.update
//
// 增量方式修改天猫商品的API。只要是此接口支持增量修改的字段，可以同时更新。（感谢爱慕旗舰店提供API命名）
type TmallitemschemaincrementupdateAPIRequest struct {
	model.Params
	// 根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明
	_xmlData string
	// 需要编辑的商品ID
	_itemId int64
}

// NewTmallitemschemaincrementupdateRequest 初始化TmallitemschemaincrementupdateAPIRequest对象
func NewTmallitemschemaincrementupdateRequest() *TmallitemschemaincrementupdateAPIRequest {
	return &TmallitemschemaincrementupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemschemaincrementupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.schema.increment.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemschemaincrementupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemschemaincrementupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXmlData is XmlData Setter
// 根据tmall.item.increment.update.schema.get生成的商品增量编辑规则入参数据。需要更新的字段，一定要在入参的XML重点update_fields字段中明确指明
func (r *TmallitemschemaincrementupdateAPIRequest) SetXmlData(_xmlData string) error {
	r._xmlData = _xmlData
	r.Set("xml_data", _xmlData)
	return nil
}

// GetXmlData XmlData Getter
func (r TmallitemschemaincrementupdateAPIRequest) GetXmlData() string {
	return r._xmlData
}

// SetItemId is ItemId Setter
// 需要编辑的商品ID
func (r *TmallitemschemaincrementupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemschemaincrementupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}
