package moscm

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosGoodsSynchinventorybycountingAPIRequest) Reset() {
	r._countingItemDto = r._countingItemDto[:0]
	r._paramCountingInfoDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSynchinventorybycountingAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.synchinventorybycounting"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSynchinventorybycountingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosGoodsSynchinventorybycountingAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaMosGoodsSynchinventorybycountingAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosGoodsSynchinventorybycountingRequest()
	},
}

// GetAlibabaMosGoodsSynchinventorybycountingRequest 从 sync.Pool 获取 AlibabaMosGoodsSynchinventorybycountingAPIRequest
func GetAlibabaMosGoodsSynchinventorybycountingAPIRequest() *AlibabaMosGoodsSynchinventorybycountingAPIRequest {
	return poolAlibabaMosGoodsSynchinventorybycountingAPIRequest.Get().(*AlibabaMosGoodsSynchinventorybycountingAPIRequest)
}

// ReleaseAlibabaMosGoodsSynchinventorybycountingAPIRequest 将 AlibabaMosGoodsSynchinventorybycountingAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosGoodsSynchinventorybycountingAPIRequest(v *AlibabaMosGoodsSynchinventorybycountingAPIRequest) {
	v.Reset()
	poolAlibabaMosGoodsSynchinventorybycountingAPIRequest.Put(v)
}
