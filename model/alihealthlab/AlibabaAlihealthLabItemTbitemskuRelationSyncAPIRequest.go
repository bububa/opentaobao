package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 API请求
// alibaba.alihealth.lab.item.tbitemsku.relation.sync
//
// 阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步
type AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest struct {
	model.Params
	// EFFECTIVE 有效，INVALID 无效
	_isvRelationStatus string
	// 检验检测项目isv侧code
	_isvItemCode string
	// 关联的淘宝商品SKU id，在商品没有sku的情况下传0
	_tbSkuId int64
	// 关联的淘宝商品 id
	_tbItemId int64
}

// NewAlibabaalihealthlabitemtbitemskurelationsyncRequest 初始化AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest对象
func NewAlibabaalihealthlabitemtbitemskurelationsyncRequest() *AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest {
	return &AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.item.tbitemsku.relation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvRelationStatus is IsvRelationStatus Setter
// EFFECTIVE 有效，INVALID 无效
func (r *AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) SetIsvRelationStatus(_isvRelationStatus string) error {
	r._isvRelationStatus = _isvRelationStatus
	r.Set("isv_relation_status", _isvRelationStatus)
	return nil
}

// GetIsvRelationStatus IsvRelationStatus Getter
func (r AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) GetIsvRelationStatus() string {
	return r._isvRelationStatus
}

// SetIsvItemCode is IsvItemCode Setter
// 检验检测项目isv侧code
func (r *AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) SetIsvItemCode(_isvItemCode string) error {
	r._isvItemCode = _isvItemCode
	r.Set("isv_item_code", _isvItemCode)
	return nil
}

// GetIsvItemCode IsvItemCode Getter
func (r AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) GetIsvItemCode() string {
	return r._isvItemCode
}

// SetTbSkuId is TbSkuId Setter
// 关联的淘宝商品SKU id，在商品没有sku的情况下传0
func (r *AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) SetTbSkuId(_tbSkuId int64) error {
	r._tbSkuId = _tbSkuId
	r.Set("tb_sku_id", _tbSkuId)
	return nil
}

// GetTbSkuId TbSkuId Getter
func (r AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) GetTbSkuId() int64 {
	return r._tbSkuId
}

// SetTbItemId is TbItemId Setter
// 关联的淘宝商品 id
func (r *AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) SetTbItemId(_tbItemId int64) error {
	r._tbItemId = _tbItemId
	r.Set("tb_item_id", _tbItemId)
	return nil
}

// GetTbItemId TbItemId Getter
func (r AlibabaalihealthlabitemtbitemskurelationsyncAPIRequest) GetTbItemId() int64 {
	return r._tbItemId
}
