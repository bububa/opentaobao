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
    _bizCode   string
    // 合同编号
    _contractCode   string
    // 摊位编码
    _storeCode   string
    // 收银比例
    _payRatio   string
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
func (r *TmallNrtStallPayratioSynchronizeRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetBizCode() string {
    return r._bizCode
}
// ContractCode Setter
// 合同编号
func (r *TmallNrtStallPayratioSynchronizeRequest) SetContractCode(_contractCode string) error {
    r._contractCode = _contractCode
    r.Set("contract_code", _contractCode)
    return nil
}

// ContractCode Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetContractCode() string {
    return r._contractCode
}
// StoreCode Setter
// 摊位编码
func (r *TmallNrtStallPayratioSynchronizeRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetStoreCode() string {
    return r._storeCode
}
// PayRatio Setter
// 收银比例
func (r *TmallNrtStallPayratioSynchronizeRequest) SetPayRatio(_payRatio string) error {
    r._payRatio = _payRatio
    r.Set("pay_ratio", _payRatio)
    return nil
}

// PayRatio Getter
func (r TmallNrtStallPayratioSynchronizeRequest) GetPayRatio() string {
    return r._payRatio
}
