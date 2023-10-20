package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpaccountgetcanusebizcodeAPIRequest 获取账户可用的bizCode API请求
// taobao.universalbp.account.get.can.use.bizcode
//
// 查询账户可用场景，查询场景名称和场景bizcode的对应关系。其中bizcode在几乎所有接口的context中需要传入。
type TaobaouniversalbpaccountgetcanusebizcodeAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaouniversalbpaccountgetcanusebizcodeRequest 初始化TaobaouniversalbpaccountgetcanusebizcodeAPIRequest对象
func NewTaobaouniversalbpaccountgetcanusebizcodeRequest() *TaobaouniversalbpaccountgetcanusebizcodeAPIRequest {
	return &TaobaouniversalbpaccountgetcanusebizcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaouniversalbpaccountgetcanusebizcodeAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.account.get.can.use.bizcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaouniversalbpaccountgetcanusebizcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaouniversalbpaccountgetcanusebizcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaouniversalbpaccountgetcanusebizcodeAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaouniversalbpaccountgetcanusebizcodeAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}
