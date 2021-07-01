package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest
阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 API请求
alibaba.alihealth.lab.item.tbitemsku.relation.sync

阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 */
type AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest struct {
	model.Params
	// EFFECTIVE 有效，INVALID 无效
	_isvRelationStatus string
	// 关联的淘宝商品SKU id，在商品没有sku的情况下传0
	_tbSkuId int64
	// 关联的淘宝商品 id
	_tbItemId int64
	// 检验检测项目isv侧code
	_isvItemCode string
}

// NewAlibabaAlihealthLabItemTbitemskuRelationSyncRequest 初始化AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest对象
func NewAlibabaAlihealthLabItemTbitemskuRelationSyncRequest() *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest {
	return &AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.item.tbitemsku.relation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IsvRelationStatus Setter
// EFFECTIVE 有效，INVALID 无效
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) SetIsvRelationStatus(_isvRelationStatus string) error {
	r._isvRelationStatus = _isvRelationStatus
	r.Set("isv_relation_status", _isvRelationStatus)
	return nil
}

// Get IsvRelationStatus Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) GetIsvRelationStatus() string {
	return r._isvRelationStatus
}

// Set is TbSkuId Setter
// 关联的淘宝商品SKU id，在商品没有sku的情况下传0
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) SetTbSkuId(_tbSkuId int64) error {
	r._tbSkuId = _tbSkuId
	r.Set("tb_sku_id", _tbSkuId)
	return nil
}

// Get TbSkuId Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) GetTbSkuId() int64 {
	return r._tbSkuId
}

// Set is TbItemId Setter
// 关联的淘宝商品 id
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) SetTbItemId(_tbItemId int64) error {
	r._tbItemId = _tbItemId
	r.Set("tb_item_id", _tbItemId)
	return nil
}

// Get TbItemId Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) GetTbItemId() int64 {
	return r._tbItemId
}

// Set is IsvItemCode Setter
// 检验检测项目isv侧code
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) SetIsvItemCode(_isvItemCode string) error {
	r._isvItemCode = _isvItemCode
	r.Set("isv_item_code", _isvItemCode)
	return nil
}

// Get IsvItemCode Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest) GetIsvItemCode() string {
	return r._isvItemCode
}
