package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlipayxiaodaiuserpermitAPIRequest 阿里金融用户授权 API请求
// alipay.xiaodai.user.permit
//
// 阿里金融为用户开通消息通道接口
type AlipayxiaodaiuserpermitAPIRequest struct {
	model.Params
	// 用户数字ID
	_userId int64
}

// NewAlipayxiaodaiuserpermitRequest 初始化AlipayxiaodaiuserpermitAPIRequest对象
func NewAlipayxiaodaiuserpermitRequest() *AlipayxiaodaiuserpermitAPIRequest {
	return &AlipayxiaodaiuserpermitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipayxiaodaiuserpermitAPIRequest) GetApiMethodName() string {
	return "alipay.xiaodai.user.permit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipayxiaodaiuserpermitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipayxiaodaiuserpermitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户数字ID
func (r *AlipayxiaodaiuserpermitAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlipayxiaodaiuserpermitAPIRequest) GetUserId() int64 {
	return r._userId
}
