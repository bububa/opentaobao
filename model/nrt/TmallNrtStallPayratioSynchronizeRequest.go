package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步摊位收银比例 API请求
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

// 初始化TmallNrtStallPayratioSynchronizeRequest对象
func NewTmallNrtStallPayratioSynchronizeRequest() *TmallNrtStallPayratioSynchronizeRequest{
    return &TmallNrtStallPayratioSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtStallPayratioSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.stall.payratio.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtStallPayratioSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizCode Setter
// 业务编码
func (r *TmallNrtStallPayratioSynchronizeRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

// BizCode Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetBizCode() string {
    return r.bizCode
}
// ContractCode Setter
// 合同编号
func (r *TmallNrtStallPayratioSynchronizeRequest) SetContractCode(contractCode string) error {
    r.contractCode = contractCode
    r.Set("contract_code", contractCode)
    return nil
}

// ContractCode Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetContractCode() string {
    return r.contractCode
}
// StoreCode Setter
// 摊位编码
func (r *TmallNrtStallPayratioSynchronizeRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

// StoreCode Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetStoreCode() string {
    return r.storeCode
}
// PayRatio Setter
// 收银比例
func (r *TmallNrtStallPayratioSynchronizeRequest) SetPayRatio(payRatio string) error {
    r.payRatio = payRatio
    r.Set("pay_ratio", payRatio)
    return nil
}

// PayRatio Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetPayRatio() string {
    return r.payRatio
}
