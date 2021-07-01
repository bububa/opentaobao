package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountListAPIRequest
OpenAccount账号信息查询 API请求
taobao.open.account.list

OpenAccount账号信息查询 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountListAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpenAccountIds Setter
// Open Account的id列表, 每次最多查询 20 个帐户
func (r *TaobaoOpenAccountListAPIRequest) SetOpenAccountIds(_openAccountIds []int64) error {
	r._openAccountIds = _openAccountIds
	r.Set("open_account_ids", _openAccountIds)
	return nil
}

// Get OpenAccountIds Getter
func (r TaobaoOpenAccountListAPIRequest) GetOpenAccountIds() []int64 {
	return r._openAccountIds
}

// Set is IsvAccountIds Setter
// ISV自己账号的id列表，isvAccountId和openAccountId二选一必填, 每次最多查询 20 个帐户
func (r *TaobaoOpenAccountListAPIRequest) SetIsvAccountIds(_isvAccountIds []string) error {
	r._isvAccountIds = _isvAccountIds
	r.Set("isv_account_ids", _isvAccountIds)
	return nil
}

// Get IsvAccountIds Getter
func (r TaobaoOpenAccountListAPIRequest) GetIsvAccountIds() []string {
	return r._isvAccountIds
}
