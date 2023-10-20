package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountDeleteAPIRequest OpenAccount删除数据 API请求
// taobao.open.account.delete
//
// OpenAccount删除数据
type TaobaoOpenAccountDeleteAPIRequest struct {
	model.Params
	// Open Account的id列表
	_openAccountIds []int64
	// ISV自己账号的id列表
	_isvAccountIds []string
}

// NewTaobaoOpenAccountDeleteRequest 初始化TaobaoOpenAccountDeleteAPIRequest对象
func NewTaobaoOpenAccountDeleteRequest() *TaobaoOpenAccountDeleteAPIRequest {
	return &TaobaoOpenAccountDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenAccountDeleteAPIRequest) Reset() {
	r._openAccountIds = r._openAccountIds[:0]
	r._isvAccountIds = r._isvAccountIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenAccountDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAccountIds is OpenAccountIds Setter
// Open Account的id列表
func (r *TaobaoOpenAccountDeleteAPIRequest) SetOpenAccountIds(_openAccountIds []int64) error {
	r._openAccountIds = _openAccountIds
	r.Set("open_account_ids", _openAccountIds)
	return nil
}

// GetOpenAccountIds OpenAccountIds Getter
func (r TaobaoOpenAccountDeleteAPIRequest) GetOpenAccountIds() []int64 {
	return r._openAccountIds
}

// SetIsvAccountIds is IsvAccountIds Setter
// ISV自己账号的id列表
func (r *TaobaoOpenAccountDeleteAPIRequest) SetIsvAccountIds(_isvAccountIds []string) error {
	r._isvAccountIds = _isvAccountIds
	r.Set("isv_account_ids", _isvAccountIds)
	return nil
}

// GetIsvAccountIds IsvAccountIds Getter
func (r TaobaoOpenAccountDeleteAPIRequest) GetIsvAccountIds() []string {
	return r._isvAccountIds
}

var poolTaobaoOpenAccountDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenAccountDeleteRequest()
	},
}

// GetTaobaoOpenAccountDeleteRequest 从 sync.Pool 获取 TaobaoOpenAccountDeleteAPIRequest
func GetTaobaoOpenAccountDeleteAPIRequest() *TaobaoOpenAccountDeleteAPIRequest {
	return poolTaobaoOpenAccountDeleteAPIRequest.Get().(*TaobaoOpenAccountDeleteAPIRequest)
}

// ReleaseTaobaoOpenAccountDeleteAPIRequest 将 TaobaoOpenAccountDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenAccountDeleteAPIRequest(v *TaobaoOpenAccountDeleteAPIRequest) {
	v.Reset()
	poolTaobaoOpenAccountDeleteAPIRequest.Put(v)
}
