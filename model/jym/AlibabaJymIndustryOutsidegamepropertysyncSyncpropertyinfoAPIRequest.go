package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest 外部上报游戏属性信息 API请求
// alibaba.jym.industry.outsidegamepropertysync.syncpropertyinfo
//
// 外部上报游戏属性信息
type AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest struct {
	model.Params
	// 属性信息DTO
	_outSideSyncGamePropertyRequestDto *OutSideSyncGamePropertyRequestDto
}

// NewAlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoRequest 初始化AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest对象
func NewAlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoRequest() *AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest {
	return &AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.outsidegamepropertysync.syncpropertyinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutSideSyncGamePropertyRequestDto is OutSideSyncGamePropertyRequestDto Setter
// 属性信息DTO
func (r *AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest) SetOutSideSyncGamePropertyRequestDto(_outSideSyncGamePropertyRequestDto *OutSideSyncGamePropertyRequestDto) error {
	r._outSideSyncGamePropertyRequestDto = _outSideSyncGamePropertyRequestDto
	r.Set("out_side_sync_game_property_request_dto", _outSideSyncGamePropertyRequestDto)
	return nil
}

// GetOutSideSyncGamePropertyRequestDto OutSideSyncGamePropertyRequestDto Getter
func (r AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest) GetOutSideSyncGamePropertyRequestDto() *OutSideSyncGamePropertyRequestDto {
	return r._outSideSyncGamePropertyRequestDto
}
