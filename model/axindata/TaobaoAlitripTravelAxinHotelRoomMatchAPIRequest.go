package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelroommatchAPIRequest 阿信酒店房型匹配 API请求
// taobao.alitrip.travel.axin.hotel.room.match
//
// 阿信酒店房型匹配
type TaobaoalitriptravelaxinhotelroommatchAPIRequest struct {
	model.Params
	// 参数
	_req *MatchedRoomRequestDto
}

// NewTaobaoalitriptravelaxinhotelroommatchRequest 初始化TaobaoalitriptravelaxinhotelroommatchAPIRequest对象
func NewTaobaoalitriptravelaxinhotelroommatchRequest() *TaobaoalitriptravelaxinhotelroommatchAPIRequest {
	return &TaobaoalitriptravelaxinhotelroommatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelroommatchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.room.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelroommatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelroommatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 参数
func (r *TaobaoalitriptravelaxinhotelroommatchAPIRequest) SetReq(_req *MatchedRoomRequestDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaoalitriptravelaxinhotelroommatchAPIRequest) GetReq() *MatchedRoomRequestDto {
	return r._req
}
