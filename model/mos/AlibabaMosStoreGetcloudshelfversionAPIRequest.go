package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosStoreGetcloudshelfversionAPIRequest
获取云货架版本信息 API请求
alibaba.mos.store.getcloudshelfversion

根据屏编号获取云货架版本信息 */
type AlibabaMosStoreGetcloudshelfversionAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
}

// NewAlibabaMosStoreGetcloudshelfversionRequest 初始化AlibabaMosStoreGetcloudshelfversionAPIRequest对象
func NewAlibabaMosStoreGetcloudshelfversionRequest() *AlibabaMosStoreGetcloudshelfversionAPIRequest {
	return &AlibabaMosStoreGetcloudshelfversionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosStoreGetcloudshelfversionAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.getcloudshelfversion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosStoreGetcloudshelfversionAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ScreenNo Setter
// 屏编号
func (r *AlibabaMosStoreGetcloudshelfversionAPIRequest) SetScreenNo(_screenNo string) error {
	r._screenNo = _screenNo
	r.Set("screen_no", _screenNo)
	return nil
}

// Get ScreenNo Getter
func (r AlibabaMosStoreGetcloudshelfversionAPIRequest) GetScreenNo() string {
	return r._screenNo
}
