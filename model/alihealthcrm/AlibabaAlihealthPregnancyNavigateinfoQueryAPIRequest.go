package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpregnancynavigateinfoqueryAPIRequest 查询底导数据 API请求
// alibaba.alihealth.pregnancy.navigateinfo.query
//
// 备孕管理--获取底部导航信息
type AlibabaalihealthpregnancynavigateinfoqueryAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
}

// NewAlibabaalihealthpregnancynavigateinfoqueryRequest 初始化AlibabaalihealthpregnancynavigateinfoqueryAPIRequest对象
func NewAlibabaalihealthpregnancynavigateinfoqueryRequest() *AlibabaalihealthpregnancynavigateinfoqueryAPIRequest {
	return &AlibabaalihealthpregnancynavigateinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpregnancynavigateinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.navigateinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpregnancynavigateinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpregnancynavigateinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaalihealthpregnancynavigateinfoqueryAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaalihealthpregnancynavigateinfoqueryAPIRequest) GetUserId() int64 {
	return r._userId
}
