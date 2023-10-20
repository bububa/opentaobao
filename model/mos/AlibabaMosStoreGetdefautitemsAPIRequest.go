package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosstoregetdefautitemsAPIRequest 获取默认状态下商品列表 API请求
// alibaba.mos.store.getdefautitems
//
// 获取默认状态下商品列表
type AlibabamosstoregetdefautitemsAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
	// 分页查询开始index
	_start int64
	// 分页查询每页记录数
	_limitCount int64
}

// NewAlibabamosstoregetdefautitemsRequest 初始化AlibabamosstoregetdefautitemsAPIRequest对象
func NewAlibabamosstoregetdefautitemsRequest() *AlibabamosstoregetdefautitemsAPIRequest {
	return &AlibabamosstoregetdefautitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosstoregetdefautitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.store.getdefautitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosstoregetdefautitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosstoregetdefautitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScreenNo is ScreenNo Setter
// 屏编号
func (r *AlibabamosstoregetdefautitemsAPIRequest) SetScreenNo(_screenNo string) error {
	r._screenNo = _screenNo
	r.Set("screen_no", _screenNo)
	return nil
}

// GetScreenNo ScreenNo Getter
func (r AlibabamosstoregetdefautitemsAPIRequest) GetScreenNo() string {
	return r._screenNo
}

// SetStart is Start Setter
// 分页查询开始index
func (r *AlibabamosstoregetdefautitemsAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r AlibabamosstoregetdefautitemsAPIRequest) GetStart() int64 {
	return r._start
}

// SetLimitCount is LimitCount Setter
// 分页查询每页记录数
func (r *AlibabamosstoregetdefautitemsAPIRequest) SetLimitCount(_limitCount int64) error {
	r._limitCount = _limitCount
	r.Set("limit_count", _limitCount)
	return nil
}

// GetLimitCount LimitCount Getter
func (r AlibabamosstoregetdefautitemsAPIRequest) GetLimitCount() int64 {
	return r._limitCount
}
