package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsSynchinventorybycountingAPIRequest 以盘点方式调整库存：传入商品实际库存 API请求
// alibaba.mos.goods.synchinventorybycounting
//
// 以盘点方式调整库存：传入商品实际库存
// 盘点单自动判断数量增减
type AlibabaMosGoodsSynchinventorybycountingAPIRequest struct {
	model.Params
	// 盘点库存项（最大列表长度：20）
	_countingItemDto []CountingItemDto
	// 专柜信息
	_paramCountingInfoDto *CountingInfoDto
}

// NewAlibabaMosGoodsSynchinventorybycountingRequest 初始化AlibabaMosGoodsSynchinventorybycountingAPIRequest对象
func NewAlibabaMosGoodsSynchinventorybycountingRequest() *AlibabaMosGoodsSynchinventorybycountingAPIRequest {
	return &AlibabaMosGoodsSynchinventorybycountingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSynchinventorybycountingAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.synchinventorybycounting"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSynchinventorybycountingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCountingItemDto is CountingItemDto Setter
// 盘点库存项（最大列表长度：20）
func (r *AlibabaMosGoodsSynchinventorybycountingAPIRequest) SetCountingItemDto(_countingItemDto []CountingItemDto) error {
	r._countingItemDto = _countingItemDto
	r.Set("counting_item_dto", _countingItemDto)
	return nil
}

// GetCountingItemDto CountingItemDto Getter
func (r AlibabaMosGoodsSynchinventorybycountingAPIRequest) GetCountingItemDto() []CountingItemDto {
	return r._countingItemDto
}

// SetParamCountingInfoDto is ParamCountingInfoDto Setter
// 专柜信息
func (r *AlibabaMosGoodsSynchinventorybycountingAPIRequest) SetParamCountingInfoDto(_paramCountingInfoDto *CountingInfoDto) error {
	r._paramCountingInfoDto = _paramCountingInfoDto
	r.Set("param_counting_info_dto", _paramCountingInfoDto)
	return nil
}

// GetParamCountingInfoDto ParamCountingInfoDto Getter
func (r AlibabaMosGoodsSynchinventorybycountingAPIRequest) GetParamCountingInfoDto() *CountingInfoDto {
	return r._paramCountingInfoDto
}
