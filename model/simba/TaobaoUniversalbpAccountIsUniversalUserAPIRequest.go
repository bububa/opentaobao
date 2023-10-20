package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpaccountisuniversaluserAPIRequest 判断用户是否迁移新bp API请求
// taobao.universalbp.account.is.universal.user
//
// 获取客户是否应使用新接口。对于迁移了新bp的客户，使用新接口，没有迁移的，使用老bp接口。不可错乱使用。
type TaobaouniversalbpaccountisuniversaluserAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbpaccountisuniversaluserRequest 初始化TaobaouniversalbpaccountisuniversaluserAPIRequest对象
func NewTaobaouniversalbpaccountisuniversaluserRequest() *TaobaouniversalbpaccountisuniversaluserAPIRequest {
	return &TaobaouniversalbpaccountisuniversaluserAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpaccountisuniversaluserAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.account.is.universal.user"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpaccountisuniversaluserAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpaccountisuniversaluserAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpaccountisuniversaluserAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpaccountisuniversaluserAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
