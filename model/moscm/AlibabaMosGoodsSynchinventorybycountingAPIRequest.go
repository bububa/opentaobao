package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosgoodssynchinventorybycountingAPIRequest 以盘点方式调整库存：传入商品实际库存 API请求
// alibaba.mos.goods.synchinventorybycounting
//
// 以盘点方式调整库存：传入商品实际库存
// 盘点单自动判断数量增减
type AlibabamosgoodssynchinventorybycountingAPIRequest struct {
	model.Params
	// 盘点库存项（最大列表长度：20）
	_countingItemDto []CountingItemDto
	// 专柜信息
	_paramCountingInfoDto *CountingInfoDto
}

// NewAlibabamosgoodssynchinventorybycountingRequest 初始化AlibabamosgoodssynchinventorybycountingAPIRequest对象
func NewAlibabamosgoodssynchinventorybycountingRequest() *AlibabamosgoodssynchinventorybycountingAPIRequest {
	return &AlibabamosgoodssynchinventorybycountingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosgoodssynchinventorybycountingAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.synchinventorybycounting"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosgoodssynchinventorybycountingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosgoodssynchinventorybycountingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCountingItemDto is CountingItemDto Setter
// 盘点库存项（最大列表长度：20）
func (r *AlibabamosgoodssynchinventorybycountingAPIRequest) SetCountingItemDto(_countingItemDto []CountingItemDto) error {
	r._countingItemDto = _countingItemDto
	r.Set("counting_item_dto", _countingItemDto)
	return nil
}

// GetCountingItemDto CountingItemDto Getter
func (r AlibabamosgoodssynchinventorybycountingAPIRequest) GetCountingItemDto() []CountingItemDto {
	return r._countingItemDto
}

// SetParamCountingInfoDto is ParamCountingInfoDto Setter
// 专柜信息
func (r *AlibabamosgoodssynchinventorybycountingAPIRequest) SetParamCountingInfoDto(_paramCountingInfoDto *CountingInfoDto) error {
	r._paramCountingInfoDto = _paramCountingInfoDto
	r.Set("param_counting_info_dto", _paramCountingInfoDto)
	return nil
}

// GetParamCountingInfoDto ParamCountingInfoDto Getter
func (r AlibabamosgoodssynchinventorybycountingAPIRequest) GetParamCountingInfoDto() *CountingInfoDto {
	return r._paramCountingInfoDto
}
