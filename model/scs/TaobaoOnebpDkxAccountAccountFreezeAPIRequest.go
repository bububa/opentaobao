package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxaccountaccountfreezeAPIRequest 创建计划后支付 API请求
// taobao.onebp.dkx.account.account.freeze
//
// 创建计划后支付。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxaccountaccountfreezeAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 账户信息参数
	_accountInfo *AccountTopDto
}

// NewTaobaoonebpdkxaccountaccountfreezeRequest 初始化TaobaoonebpdkxaccountaccountfreezeAPIRequest对象
func NewTaobaoonebpdkxaccountaccountfreezeRequest() *TaobaoonebpdkxaccountaccountfreezeAPIRequest {
	return &TaobaoonebpdkxaccountaccountfreezeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoonebpdkxaccountaccountfreezeAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.account.account.freeze"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoonebpdkxaccountaccountfreezeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoonebpdkxaccountaccountfreezeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoonebpdkxaccountaccountfreezeAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoonebpdkxaccountaccountfreezeAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetAccountInfo is AccountInfo Setter
// 账户信息参数
func (r *TaobaoonebpdkxaccountaccountfreezeAPIRequest) SetAccountInfo(_accountInfo *AccountTopDto) error {
	r._accountInfo = _accountInfo
	r.Set("account_info", _accountInfo)
	return nil
}

// GetAccountInfo AccountInfo Getter
func (r TaobaoonebpdkxaccountaccountfreezeAPIRequest) GetAccountInfo() *AccountTopDto {
	return r._accountInfo
}
