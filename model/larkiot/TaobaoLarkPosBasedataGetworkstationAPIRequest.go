package larkiot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkPosBasedataGetworkstationAPIRequest 根据影城id工作站和macId获取工作站 API请求
// taobao.lark.pos.basedata.getworkstation
//
// 获取单独工作站
type TaobaoLarkPosBasedataGetworkstationAPIRequest struct {
	model.Params
	// 影城cinemaLinkId
	_cinemaLinkId string
	// 终端编码
	_posCode string
}

// NewTaobaoLarkPosBasedataGetworkstationRequest 初始化TaobaoLarkPosBasedataGetworkstationAPIRequest对象
func NewTaobaoLarkPosBasedataGetworkstationRequest() *TaobaoLarkPosBasedataGetworkstationAPIRequest {
	return &TaobaoLarkPosBasedataGetworkstationAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLarkPosBasedataGetworkstationAPIRequest) Reset() {
	r._cinemaLinkId = ""
	r._posCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLarkPosBasedataGetworkstationAPIRequest) GetApiMethodName() string {
	return "taobao.lark.pos.basedata.getworkstation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLarkPosBasedataGetworkstationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLarkPosBasedataGetworkstationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCinemaLinkId is CinemaLinkId Setter
// 影城cinemaLinkId
func (r *TaobaoLarkPosBasedataGetworkstationAPIRequest) SetCinemaLinkId(_cinemaLinkId string) error {
	r._cinemaLinkId = _cinemaLinkId
	r.Set("cinema_link_id", _cinemaLinkId)
	return nil
}

// GetCinemaLinkId CinemaLinkId Getter
func (r TaobaoLarkPosBasedataGetworkstationAPIRequest) GetCinemaLinkId() string {
	return r._cinemaLinkId
}

// SetPosCode is PosCode Setter
// 终端编码
func (r *TaobaoLarkPosBasedataGetworkstationAPIRequest) SetPosCode(_posCode string) error {
	r._posCode = _posCode
	r.Set("pos_code", _posCode)
	return nil
}

// GetPosCode PosCode Getter
func (r TaobaoLarkPosBasedataGetworkstationAPIRequest) GetPosCode() string {
	return r._posCode
}

var poolTaobaoLarkPosBasedataGetworkstationAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLarkPosBasedataGetworkstationRequest()
	},
}

// GetTaobaoLarkPosBasedataGetworkstationRequest 从 sync.Pool 获取 TaobaoLarkPosBasedataGetworkstationAPIRequest
func GetTaobaoLarkPosBasedataGetworkstationAPIRequest() *TaobaoLarkPosBasedataGetworkstationAPIRequest {
	return poolTaobaoLarkPosBasedataGetworkstationAPIRequest.Get().(*TaobaoLarkPosBasedataGetworkstationAPIRequest)
}

// ReleaseTaobaoLarkPosBasedataGetworkstationAPIRequest 将 TaobaoLarkPosBasedataGetworkstationAPIRequest 放入 sync.Pool
func ReleaseTaobaoLarkPosBasedataGetworkstationAPIRequest(v *TaobaoLarkPosBasedataGetworkstationAPIRequest) {
	v.Reset()
	poolTaobaoLarkPosBasedataGetworkstationAPIRequest.Put(v)
}
