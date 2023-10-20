package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelbookinfoqueryAPIRequest 飞猪度假-订单二次预约查询接口 API请求
// alitrip.travel.bookinfo.query
//
// 飞猪度假订单二次预约详情查询接口
type AlitriptravelbookinfoqueryAPIRequest struct {
	model.Params
	// 预定信息id
	_bookinfoId int64
}

// NewAlitriptravelbookinfoqueryRequest 初始化AlitriptravelbookinfoqueryAPIRequest对象
func NewAlitriptravelbookinfoqueryRequest() *AlitriptravelbookinfoqueryAPIRequest {
	return &AlitriptravelbookinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelbookinfoqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.bookinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelbookinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelbookinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBookinfoId is BookinfoId Setter
// 预定信息id
func (r *AlitriptravelbookinfoqueryAPIRequest) SetBookinfoId(_bookinfoId int64) error {
	r._bookinfoId = _bookinfoId
	r.Set("bookinfo_id", _bookinfoId)
	return nil
}

// GetBookinfoId BookinfoId Getter
func (r AlitriptravelbookinfoqueryAPIRequest) GetBookinfoId() int64 {
	return r._bookinfoId
}
