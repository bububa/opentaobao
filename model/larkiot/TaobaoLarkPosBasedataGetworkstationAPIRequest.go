package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolarkposbasedatagetworkstationAPIRequest 根据影城id工作站和macId获取工作站 API请求
// taobao.lark.pos.basedata.getworkstation
//
// 获取单独工作站
type TaobaolarkposbasedatagetworkstationAPIRequest struct {
	model.Params
	// 影城cinemaLinkId
	_cinemaLinkId string
	// 终端编码
	_posCode string
}

// NewTaobaolarkposbasedatagetworkstationRequest 初始化TaobaolarkposbasedatagetworkstationAPIRequest对象
func NewTaobaolarkposbasedatagetworkstationRequest() *TaobaolarkposbasedatagetworkstationAPIRequest {
	return &TaobaolarkposbasedatagetworkstationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolarkposbasedatagetworkstationAPIRequest) GetApiMethodName() string {
	return "taobao.lark.pos.basedata.getworkstation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolarkposbasedatagetworkstationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolarkposbasedatagetworkstationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCinemaLinkId is CinemaLinkId Setter
// 影城cinemaLinkId
func (r *TaobaolarkposbasedatagetworkstationAPIRequest) SetCinemaLinkId(_cinemaLinkId string) error {
	r._cinemaLinkId = _cinemaLinkId
	r.Set("cinema_link_id", _cinemaLinkId)
	return nil
}

// GetCinemaLinkId CinemaLinkId Getter
func (r TaobaolarkposbasedatagetworkstationAPIRequest) GetCinemaLinkId() string {
	return r._cinemaLinkId
}

// SetPosCode is PosCode Setter
// 终端编码
func (r *TaobaolarkposbasedatagetworkstationAPIRequest) SetPosCode(_posCode string) error {
	r._posCode = _posCode
	r.Set("pos_code", _posCode)
	return nil
}

// GetPosCode PosCode Getter
func (r TaobaolarkposbasedatagetworkstationAPIRequest) GetPosCode() string {
	return r._posCode
}
