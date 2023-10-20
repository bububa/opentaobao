package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest 获取账户可用的bizCode API请求
// taobao.universalbp.account.get.can.use.bizcode
//
// 查询账户可用场景，查询场景名称和场景bizcode的对应关系。其中bizcode在几乎所有接口的context中需要传入。
type TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
}

// NewTaobaoUniversalbpAccountGetCanUseBizcodeRequest 初始化TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest对象
func NewTaobaoUniversalbpAccountGetCanUseBizcodeRequest() *TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest {
	return &TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest) Reset() {
	r._topServiceContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.account.get.can.use.bizcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

var poolTaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpAccountGetCanUseBizcodeRequest()
	},
}

// GetTaobaoUniversalbpAccountGetCanUseBizcodeRequest 从 sync.Pool 获取 TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest
func GetTaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest() *TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest {
	return poolTaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest.Get().(*TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest)
}

// ReleaseTaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest 将 TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest(v *TaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpAccountGetCanUseBizcodeAPIRequest.Put(v)
}
