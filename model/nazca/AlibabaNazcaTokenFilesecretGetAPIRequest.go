package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcatokenfilesecretgetAPIRequest 获取文件秘钥 API请求
// alibaba.nazca.token.filesecret.get
//
// 获取文件秘钥
type AlibabanazcatokenfilesecretgetAPIRequest struct {
	model.Params
	// 客户在1688的唯一标识
	_platformUserId string
	// 合同编号
	_contractNum string
}

// NewAlibabanazcatokenfilesecretgetRequest 初始化AlibabanazcatokenfilesecretgetAPIRequest对象
func NewAlibabanazcatokenfilesecretgetRequest() *AlibabanazcatokenfilesecretgetAPIRequest {
	return &AlibabanazcatokenfilesecretgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanazcatokenfilesecretgetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.filesecret.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanazcatokenfilesecretgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanazcatokenfilesecretgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPlatformUserId is PlatformUserId Setter
// 客户在1688的唯一标识
func (r *AlibabanazcatokenfilesecretgetAPIRequest) SetPlatformUserId(_platformUserId string) error {
	r._platformUserId = _platformUserId
	r.Set("platform_user_id", _platformUserId)
	return nil
}

// GetPlatformUserId PlatformUserId Getter
func (r AlibabanazcatokenfilesecretgetAPIRequest) GetPlatformUserId() string {
	return r._platformUserId
}

// SetContractNum is ContractNum Setter
// 合同编号
func (r *AlibabanazcatokenfilesecretgetAPIRequest) SetContractNum(_contractNum string) error {
	r._contractNum = _contractNum
	r.Set("contract_num", _contractNum)
	return nil
}

// GetContractNum ContractNum Getter
func (r AlibabanazcatokenfilesecretgetAPIRequest) GetContractNum() string {
	return r._contractNum
}
