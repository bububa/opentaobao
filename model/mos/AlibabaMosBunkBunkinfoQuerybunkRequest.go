package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据合同号查询铺位信息 APIRequest
alibaba.mos.bunk.bunkinfo.querybunk

根据合同号查询铺位信息
*/
type AlibabaMosBunkBunkinfoQuerybunkRequest struct {
    model.Params

    // 门店号
    storeNo   string 

    // 合同状态集合
    statusList   []string 

    // 合同号集合
    contractCodes   []string 

}

func NewAlibabaMosBunkBunkinfoQuerybunkRequest() *AlibabaMosBunkBunkinfoQuerybunkRequest{
    return &AlibabaMosBunkBunkinfoQuerybunkRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetApiMethodName() string {
    return "alibaba.mos.bunk.bunkinfo.querybunk"
}

func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosBunkBunkinfoQuerybunkRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetStoreNo() string {
    return r.storeNo
}

func (r *AlibabaMosBunkBunkinfoQuerybunkRequest) SetStatusList(statusList []string) error {
    r.statusList = statusList
    r.Set("status_list", statusList)
    return nil
}

func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetStatusList() []string {
    return r.statusList
}

func (r *AlibabaMosBunkBunkinfoQuerybunkRequest) SetContractCodes(contractCodes []string) error {
    r.contractCodes = contractCodes
    r.Set("contract_codes", contractCodes)
    return nil
}

func (r AlibabaMosBunkBunkinfoQuerybunkRequest) GetContractCodes() []string {
    return r.contractCodes
}

