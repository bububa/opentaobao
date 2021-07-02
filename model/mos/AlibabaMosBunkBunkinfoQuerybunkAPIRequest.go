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
	// 门店号
	_storeNo string
	// 合同状态集合
	_statusList []string
	// 合同号集合
	_contractCodes []string
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
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreNo Setter
// 门店号
func (r *AlibabaMosBunkBunkinfoQuerybunkAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// Get StoreNo Getter
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// Set is StatusList Setter
// 合同状态集合
func (r *AlibabaMosBunkBunkinfoQuerybunkAPIRequest) SetStatusList(_statusList []string) error {
	r._statusList = _statusList
	r.Set("status_list", _statusList)
	return nil
}

// Get StatusList Getter
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetStatusList() []string {
	return r._statusList
}

// Set is ContractCodes Setter
// 合同号集合
func (r *AlibabaMosBunkBunkinfoQuerybunkAPIRequest) SetContractCodes(_contractCodes []string) error {
	r._contractCodes = _contractCodes
	r.Set("contract_codes", _contractCodes)
	return nil
}

// Get ContractCodes Getter
func (r AlibabaMosBunkBunkinfoQuerybunkAPIRequest) GetContractCodes() []string {
	return r._contractCodes
}
