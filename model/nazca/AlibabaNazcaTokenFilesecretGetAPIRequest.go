package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaTokenFilesecretGetAPIRequest 获取文件秘钥 API请求
// alibaba.nazca.token.filesecret.get
//
// 获取文件秘钥
type AlibabaNazcaTokenFilesecretGetAPIRequest struct {
	model.Params
	// 客户在1688的唯一标识
	_platformUserId string
	// 合同编号
	_contractNum string
}

// NewAlibabaNazcaTokenFilesecretGetRequest 初始化AlibabaNazcaTokenFilesecretGetAPIRequest对象
func NewAlibabaNazcaTokenFilesecretGetRequest() *AlibabaNazcaTokenFilesecretGetAPIRequest {
	return &AlibabaNazcaTokenFilesecretGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenFilesecretGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.token.filesecret.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenFilesecretGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PlatformUserId Setter
// 客户在1688的唯一标识
func (r *AlibabaNazcaTokenFilesecretGetAPIRequest) SetPlatformUserId(_platformUserId string) error {
	r._platformUserId = _platformUserId
	r.Set("platform_user_id", _platformUserId)
	return nil
}

// Get PlatformUserId Getter
func (r AlibabaNazcaTokenFilesecretGetAPIRequest) GetPlatformUserId() string {
	return r._platformUserId
}

// Set is ContractNum Setter
// 合同编号
func (r *AlibabaNazcaTokenFilesecretGetAPIRequest) SetContractNum(_contractNum string) error {
	r._contractNum = _contractNum
	r.Set("contract_num", _contractNum)
	return nil
}

// Get ContractNum Getter
func (r AlibabaNazcaTokenFilesecretGetAPIRequest) GetContractNum() string {
	return r._contractNum
}
