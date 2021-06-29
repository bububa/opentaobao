package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据合同号查询铺位信息 API请求
alibaba.mos.bunk.bunkinfo.querybunk

根据合同号查询铺位信息
*/
type AlibabaMosBunkBunkinfoQuerybunkRequest struct {
    model.Params
    // 门店号
    _storeNo   string
    // 合同状态集合
    _statusList   []string
    // 合同号集合
    _contractCodes   []string
}

// 初始化AlibabaMosBunkBunkinfoQuerybunkRequest对象
func NewAlibabaMosBunkBunkinfoQuerybunkRequest() *AlibabaMosBunkBunkinfoQuerybunkRequest{
    return &AlibabaMosBunkBunkinfoQuerybunkRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetApiMethodName() string {
    return "alibaba.mos.bunk.bunkinfo.querybunk"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreNo Setter
// 门店号
func (r *AlibabaMosBunkBunkinfoQuerybunkRequest) SetStoreNo(_storeNo string) error {
    r._storeNo = _storeNo
    r.Set("store_no", _storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetStoreNo() string {
    return r._storeNo
}
// StatusList Setter
// 合同状态集合
func (r *AlibabaMosBunkBunkinfoQuerybunkRequest) SetStatusList(_statusList []string) error {
    r._statusList = _statusList
    r.Set("status_list", _statusList)
    return nil
}

// StatusList Getter
func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetStatusList() []string {
    return r._statusList
}
// ContractCodes Setter
// 合同号集合
func (r *AlibabaMosBunkBunkinfoQuerybunkRequest) SetContractCodes(_contractCodes []string) error {
    r._contractCodes = _contractCodes
    r.Set("contract_codes", _contractCodes)
    return nil
}

// ContractCodes Getter
func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetContractCodes() []string {
    return r._contractCodes
}
