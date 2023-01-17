package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosBunkBunkinfoQuerybunkAPIRequest 根据合同号查询铺位信息 API请求
// alibaba.mos.bunk.bunkinfo.querybunk
//
// 根据合同号查询铺位信息
type AlibabaMosBunkBunkinfoQuerybunkAPIRequest struct {
	model.Params
	// 合同状态集合
	_statusList []string
	// 合同号集合
	_contractCodes []string
	// 门店号
	_storeNo string
}

// NewAlibabaMosBunkBunkinfoQuerybunkRequest 初始化AlibabaMosBunkBunkinfoQuerybunkAPIRequest对象
func NewAlibabaMosBunkBunkinfoQuerybunkRequest() *AlibabaMosBunkBunkinfoQuerybunkAPIRequest {
	return &AlibabaMosBunkBunkinfoQuerybunkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.bunk.bunkinfo.querybunk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatusList is StatusList Setter
// 合同状态集合
func (r *AlibabaMosBunkBunkinfoQuerybunkAPIRequest) SetStatusList(_statusList []string) error {
	r._statusList = _statusList
	r.Set("status_list", _statusList)
	return nil
}

// GetStatusList StatusList Getter
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetStatusList() []string {
	return r._statusList
}

// SetContractCodes is ContractCodes Setter
// 合同号集合
func (r *AlibabaMosBunkBunkinfoQuerybunkAPIRequest) SetContractCodes(_contractCodes []string) error {
	r._contractCodes = _contractCodes
	r.Set("contract_codes", _contractCodes)
	return nil
}

// GetContractCodes ContractCodes Getter
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetContractCodes() []string {
	return r._contractCodes
}

// SetStoreNo is StoreNo Setter
// 门店号
func (r *AlibabaMosBunkBunkinfoQuerybunkAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetStoreNo() string {
	return r._storeNo
}
