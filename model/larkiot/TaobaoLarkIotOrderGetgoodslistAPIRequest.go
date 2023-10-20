package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaolarkiotordergetgoodslistAPIRequest iot渠道获取卖品信息 API请求
// taobao.lark.iot.order.getgoodslist
//
// iot无人超市服务商通过接口获取影院的可售卖品数据
type TaobaolarkiotordergetgoodslistAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
	// 影院内码
	_cinemaLinkId string
}

// NewTaobaolarkiotordergetgoodslistRequest 初始化TaobaolarkiotordergetgoodslistAPIRequest对象
func NewTaobaolarkiotordergetgoodslistRequest() *TaobaolarkiotordergetgoodslistAPIRequest {
	return &TaobaolarkiotordergetgoodslistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaolarkiotordergetgoodslistAPIRequest) GetApiMethodName() string {
	return "taobao.lark.iot.order.getgoodslist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaolarkiotordergetgoodslistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaolarkiotordergetgoodslistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *TaobaolarkiotordergetgoodslistAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TaobaolarkiotordergetgoodslistAPIRequest) GetChannelCode() string {
	return r._channelCode
}

// SetCinemaLinkId is CinemaLinkId Setter
// 影院内码
func (r *TaobaolarkiotordergetgoodslistAPIRequest) SetCinemaLinkId(_cinemaLinkId string) error {
	r._cinemaLinkId = _cinemaLinkId
	r.Set("cinema_link_id", _cinemaLinkId)
	return nil
}

// GetCinemaLinkId CinemaLinkId Getter
func (r TaobaolarkiotordergetgoodslistAPIRequest) GetCinemaLinkId() string {
	return r._cinemaLinkId
}
