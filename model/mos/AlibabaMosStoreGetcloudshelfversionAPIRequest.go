package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosstoregetcloudshelfversionAPIRequest 获取云货架版本信息 API请求
// alibaba.mos.store.getcloudshelfversion
//
// 根据屏编号获取云货架版本信息
type AlibabamosstoregetcloudshelfversionAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
}

// NewAlibabamosstoregetcloudshelfversionRequest 初始化AlibabamosstoregetcloudshelfversionAPIRequest对象
func NewAlibabamosstoregetcloudshelfversionRequest() *AlibabamosstoregetcloudshelfversionAPIRequest {
	return &AlibabamosstoregetcloudshelfversionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosstoregetcloudshelfversionAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.getcloudshelfversion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosstoregetcloudshelfversionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosstoregetcloudshelfversionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreenNo is ScreenNo Setter
// 屏编号
func (r *AlibabamosstoregetcloudshelfversionAPIRequest) SetScreenNo(_screenNo string) error {
	r._screenNo = _screenNo
	r.Set("screen_no", _screenNo)
	return nil
}

// GetScreenNo ScreenNo Getter
func (r AlibabamosstoregetcloudshelfversionAPIRequest) GetScreenNo() string {
	return r._screenNo
}
