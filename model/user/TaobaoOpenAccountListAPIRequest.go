package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountListAPIRequest OpenAccount账号信息查询 API请求
// taobao.open.account.list
//
// OpenAccount账号信息查询
type TaobaoOpenAccountListAPIRequest struct {
	model.Params
	// Open Account的id列表, 每次最多查询 20 个帐户
	_openAccountIds []int64
	// ISV自己账号的id列表，isvAccountId和openAccountId二选一必填, 每次最多查询 20 个帐户
	_isvAccountIds []string
}

// NewTaobaoOpenAccountListRequest 初始化TaobaoOpenAccountListAPIRequest对象
func NewTaobaoOpenAccountListRequest() *TaobaoOpenAccountListAPIRequest {
	return &TaobaoOpenAccountListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenAccountListAPIRequest) Reset() {
	r._openAccountIds = r._openAccountIds[:0]
	r._isvAccountIds = r._isvAccountIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountListAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenAccountListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAccountIds is OpenAccountIds Setter
// Open Account的id列表, 每次最多查询 20 个帐户
func (r *TaobaoOpenAccountListAPIRequest) SetOpenAccountIds(_openAccountIds []int64) error {
	r._openAccountIds = _openAccountIds
	r.Set("open_account_ids", _openAccountIds)
	return nil
}

// GetOpenAccountIds OpenAccountIds Getter
func (r TaobaoOpenAccountListAPIRequest) GetOpenAccountIds() []int64 {
	return r._openAccountIds
}

// SetIsvAccountIds is IsvAccountIds Setter
// ISV自己账号的id列表，isvAccountId和openAccountId二选一必填, 每次最多查询 20 个帐户
func (r *TaobaoOpenAccountListAPIRequest) SetIsvAccountIds(_isvAccountIds []string) error {
	r._isvAccountIds = _isvAccountIds
	r.Set("isv_account_ids", _isvAccountIds)
	return nil
}

// GetIsvAccountIds IsvAccountIds Getter
func (r TaobaoOpenAccountListAPIRequest) GetIsvAccountIds() []string {
	return r._isvAccountIds
}

var poolTaobaoOpenAccountListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenAccountListRequest()
	},
}

// GetTaobaoOpenAccountListRequest 从 sync.Pool 获取 TaobaoOpenAccountListAPIRequest
func GetTaobaoOpenAccountListAPIRequest() *TaobaoOpenAccountListAPIRequest {
	return poolTaobaoOpenAccountListAPIRequest.Get().(*TaobaoOpenAccountListAPIRequest)
}

// ReleaseTaobaoOpenAccountListAPIRequest 将 TaobaoOpenAccountListAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenAccountListAPIRequest(v *TaobaoOpenAccountListAPIRequest) {
	v.Reset()
	poolTaobaoOpenAccountListAPIRequest.Put(v)
}
