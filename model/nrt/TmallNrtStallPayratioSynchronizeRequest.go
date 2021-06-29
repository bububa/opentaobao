package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步摊位收银比例 APIRequest
tmall.nrt.stall.payratio.synchronize

ISV同步摊位收银比例到阿里
*/
type TmallNrtStallPayratioSynchronizeRequest struct {
    model.Params

    // 业务编码
    bizCode   string 

    // 合同编号
    contractCode   string 

    // 摊位编码
    storeCode   string 

    // 收银比例
    payRatio   string 

}

func NewTmallNrtStallPayratioSynchronizeRequest() *TmallNrtStallPayratioSynchronizeRequest{
    return &TmallNrtStallPayratioSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtStallPayratioSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.stall.payratio.synchronize"
}

func (r TmallNrtStallPayratioSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtStallPayratioSynchronizeRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r TmallNrtStallPayratioSynchronizeRequest) GetBizCode() string {
    return r.bizCode
}

func (r *TmallNrtStallPayratioSynchronizeRequest) SetContractCode(contractCode string) error {
    r.contractCode = contractCode
    r.Set("contract_code", contractCode)
    return nil
}

func (r TmallNrtStallPayratioSynchronizeRequest) GetContractCode() string {
    return r.contractCode
}

func (r *TmallNrtStallPayratioSynchronizeRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TmallNrtStallPayratioSynchronizeRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TmallNrtStallPayratioSynchronizeRequest) SetPayRatio(payRatio string) error {
    r.payRatio = payRatio
    r.Set("pay_ratio", payRatio)
    return nil
}

func (r TmallNrtStallPayratioSynchronizeRequest) GetPayRatio() string {
    return r.payRatio
}

