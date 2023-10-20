package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelmatchAPIRequest 酒店匹配接口-阿信 API请求
// taobao.alitrip.travel.axin.hotel.match
//
// 酒店匹配接口-阿信
type TaobaoalitriptravelaxinhotelmatchAPIRequest struct {
	model.Params
	// 请求
	_req *MatchedHotelRequestDto
}

// NewTaobaoalitriptravelaxinhotelmatchRequest 初始化TaobaoalitriptravelaxinhotelmatchAPIRequest对象
func NewTaobaoalitriptravelaxinhotelmatchRequest() *TaobaoalitriptravelaxinhotelmatchAPIRequest {
	return &TaobaoalitriptravelaxinhotelmatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelmatchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelmatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelmatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求
func (r *TaobaoalitriptravelaxinhotelmatchAPIRequest) SetReq(_req *MatchedHotelRequestDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaoalitriptravelaxinhotelmatchAPIRequest) GetReq() *MatchedHotelRequestDto {
	return r._req
}
