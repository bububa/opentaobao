package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayXiaodaiUserPermitAPIRequest 阿里金融用户授权 API请求
// alipay.xiaodai.user.permit
//
// 阿里金融为用户开通消息通道接口
type AlipayXiaodaiUserPermitAPIRequest struct {
	model.Params
	// 用户数字ID
	_userId int64
}

// NewAlipayXiaodaiUserPermitRequest 初始化AlipayXiaodaiUserPermitAPIRequest对象
func NewAlipayXiaodaiUserPermitRequest() *AlipayXiaodaiUserPermitAPIRequest {
	return &AlipayXiaodaiUserPermitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlipayXiaodaiUserPermitAPIRequest) Reset() {
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipayXiaodaiUserPermitAPIRequest) GetApiMethodName() string {
	return "alipay.xiaodai.user.permit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipayXiaodaiUserPermitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipayXiaodaiUserPermitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户数字ID
func (r *AlipayXiaodaiUserPermitAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlipayXiaodaiUserPermitAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolAlipayXiaodaiUserPermitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlipayXiaodaiUserPermitRequest()
	},
}

// GetAlipayXiaodaiUserPermitRequest 从 sync.Pool 获取 AlipayXiaodaiUserPermitAPIRequest
func GetAlipayXiaodaiUserPermitAPIRequest() *AlipayXiaodaiUserPermitAPIRequest {
	return poolAlipayXiaodaiUserPermitAPIRequest.Get().(*AlipayXiaodaiUserPermitAPIRequest)
}

// ReleaseAlipayXiaodaiUserPermitAPIRequest 将 AlipayXiaodaiUserPermitAPIRequest 放入 sync.Pool
func ReleaseAlipayXiaodaiUserPermitAPIRequest(v *AlipayXiaodaiUserPermitAPIRequest) {
	v.Reset()
	poolAlipayXiaodaiUserPermitAPIRequest.Put(v)
}
