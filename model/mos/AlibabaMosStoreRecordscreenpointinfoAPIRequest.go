package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosstorerecordscreenpointinfoAPIRequest 云屏埋点数据记录接口 API请求
// alibaba.mos.store.recordscreenpointinfo
//
// 记录云屏埋点数据
type AlibabamosstorerecordscreenpointinfoAPIRequest struct {
	model.Params
	// 云屏埋点信息
	_screenPointInfo string
}

// NewAlibabamosstorerecordscreenpointinfoRequest 初始化AlibabamosstorerecordscreenpointinfoAPIRequest对象
func NewAlibabamosstorerecordscreenpointinfoRequest() *AlibabamosstorerecordscreenpointinfoAPIRequest {
	return &AlibabamosstorerecordscreenpointinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosstorerecordscreenpointinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.recordscreenpointinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosstorerecordscreenpointinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosstorerecordscreenpointinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreenPointInfo is ScreenPointInfo Setter
// 云屏埋点信息
func (r *AlibabamosstorerecordscreenpointinfoAPIRequest) SetScreenPointInfo(_screenPointInfo string) error {
	r._screenPointInfo = _screenPointInfo
	r.Set("screen_point_info", _screenPointInfo)
	return nil
}

// GetScreenPointInfo ScreenPointInfo Getter
func (r AlibabamosstorerecordscreenpointinfoAPIRequest) GetScreenPointInfo() string {
	return r._screenPointInfo
}
