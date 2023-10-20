package larkiot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkIotOrderGetgoodslistAPIRequest iot渠道获取卖品信息 API请求
// taobao.lark.iot.order.getgoodslist
//
// iot无人超市服务商通过接口获取影院的可售卖品数据
type TaobaoLarkIotOrderGetgoodslistAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
	// 影院内码
	_cinemaLinkId string
}

// NewTaobaoLarkIotOrderGetgoodslistRequest 初始化TaobaoLarkIotOrderGetgoodslistAPIRequest对象
func NewTaobaoLarkIotOrderGetgoodslistRequest() *TaobaoLarkIotOrderGetgoodslistAPIRequest {
	return &TaobaoLarkIotOrderGetgoodslistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLarkIotOrderGetgoodslistAPIRequest) Reset() {
	r._channelCode = ""
	r._cinemaLinkId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderGetgoodslistAPIRequest) GetApiMethodName() string {
	return "taobao.lark.iot.order.getgoodslist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderGetgoodslistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLarkIotOrderGetgoodslistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderGetgoodslistAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TaobaoLarkIotOrderGetgoodslistAPIRequest) GetChannelCode() string {
	return r._channelCode
}

// SetCinemaLinkId is CinemaLinkId Setter
// 影院内码
func (r *TaobaoLarkIotOrderGetgoodslistAPIRequest) SetCinemaLinkId(_cinemaLinkId string) error {
	r._cinemaLinkId = _cinemaLinkId
	r.Set("cinema_link_id", _cinemaLinkId)
	return nil
}

// GetCinemaLinkId CinemaLinkId Getter
func (r TaobaoLarkIotOrderGetgoodslistAPIRequest) GetCinemaLinkId() string {
	return r._cinemaLinkId
}

var poolTaobaoLarkIotOrderGetgoodslistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLarkIotOrderGetgoodslistRequest()
	},
}

// GetTaobaoLarkIotOrderGetgoodslistRequest 从 sync.Pool 获取 TaobaoLarkIotOrderGetgoodslistAPIRequest
func GetTaobaoLarkIotOrderGetgoodslistAPIRequest() *TaobaoLarkIotOrderGetgoodslistAPIRequest {
	return poolTaobaoLarkIotOrderGetgoodslistAPIRequest.Get().(*TaobaoLarkIotOrderGetgoodslistAPIRequest)
}

// ReleaseTaobaoLarkIotOrderGetgoodslistAPIRequest 将 TaobaoLarkIotOrderGetgoodslistAPIRequest 放入 sync.Pool
func ReleaseTaobaoLarkIotOrderGetgoodslistAPIRequest(v *TaobaoLarkIotOrderGetgoodslistAPIRequest) {
	v.Reset()
	poolTaobaoLarkIotOrderGetgoodslistAPIRequest.Put(v)
}
