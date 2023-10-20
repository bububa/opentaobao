package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptravelsyncAPIRequest 差旅申请单同步接口 API请求
// alibaba.happytrip.travel.sync
//
// 以外部差旅申请单id（outer_travel_head_id）为主键，保存或更新差旅单信息到欢行系统中
type AlibabahappytriptravelsyncAPIRequest struct {
	model.Params
	// 差旅申请单对象
	_travelHeadDto *TravelHeadDto
}

// NewAlibabahappytriptravelsyncRequest 初始化AlibabahappytriptravelsyncAPIRequest对象
func NewAlibabahappytriptravelsyncRequest() *AlibabahappytriptravelsyncAPIRequest {
	return &AlibabahappytriptravelsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptravelsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.travel.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptravelsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptravelsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTravelHeadDto is TravelHeadDto Setter
// 差旅申请单对象
func (r *AlibabahappytriptravelsyncAPIRequest) SetTravelHeadDto(_travelHeadDto *TravelHeadDto) error {
	r._travelHeadDto = _travelHeadDto
	r.Set("travel_head_dto", _travelHeadDto)
	return nil
}

// GetTravelHeadDto TravelHeadDto Getter
func (r AlibabahappytriptravelsyncAPIRequest) GetTravelHeadDto() *TravelHeadDto {
	return r._travelHeadDto
}
