package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest 外部查询游戏属性库属性信息 API请求
// alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo
//
// 外部查询游戏属性库属性信息
type AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest struct {
	model.Params
	// 属性信息查询DTO
	_outSideQueryGamePropertyInfoRequestDto *OutSideQueryGamePropertyInfoRequestDto
}

// NewAlibabajymindustryoutsidegamepropertysyncquerypropertyinfoRequest 初始化AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest对象
func NewAlibabajymindustryoutsidegamepropertysyncquerypropertyinfoRequest() *AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest {
	return &AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutSideQueryGamePropertyInfoRequestDto is OutSideQueryGamePropertyInfoRequestDto Setter
// 属性信息查询DTO
func (r *AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest) SetOutSideQueryGamePropertyInfoRequestDto(_outSideQueryGamePropertyInfoRequestDto *OutSideQueryGamePropertyInfoRequestDto) error {
	r._outSideQueryGamePropertyInfoRequestDto = _outSideQueryGamePropertyInfoRequestDto
	r.Set("out_side_query_game_property_info_request_dto", _outSideQueryGamePropertyInfoRequestDto)
	return nil
}

// GetOutSideQueryGamePropertyInfoRequestDto OutSideQueryGamePropertyInfoRequestDto Getter
func (r AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest) GetOutSideQueryGamePropertyInfoRequestDto() *OutSideQueryGamePropertyInfoRequestDto {
	return r._outSideQueryGamePropertyInfoRequestDto
}
