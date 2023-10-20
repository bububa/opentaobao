package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreGetstorelistAPIRequest 根据屏编号获取专柜集 API请求
// alibaba.mos.store.getstorelist
//
// 根据屏编号获取专柜集
type AlibabaMosStoreGetstorelistAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
}

// NewAlibabaMosStoreGetstorelistRequest 初始化AlibabaMosStoreGetstorelistAPIRequest对象
func NewAlibabaMosStoreGetstorelistRequest() *AlibabaMosStoreGetstorelistAPIRequest {
	return &AlibabaMosStoreGetstorelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreGetstorelistAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.getstorelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreGetstorelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosStoreGetstorelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreenNo is ScreenNo Setter
// 屏编号
func (r *AlibabaMosStoreGetstorelistAPIRequest) SetScreenNo(_screenNo string) error {
	r._screenNo = _screenNo
	r.Set("screen_no", _screenNo)
	return nil
}

// GetScreenNo ScreenNo Getter
func (r AlibabaMosStoreGetstorelistAPIRequest) GetScreenNo() string {
	return r._screenNo
}
