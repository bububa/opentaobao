package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
大pos银行卡查账接口 APIResponse
alibaba.mj.oc.bigpos.banksale.query

大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
*/
type AlibabaMjOcBigposBanksaleQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMjOcBigposBanksaleQueryResponse `json:"alibaba_mj_oc_bigpos_banksale_query_response,omitempty"` 
    AlibabaMjOcBigposBanksaleQueryResponse
}

/* model for simplify = false
type AlibabaMjOcBigposBanksaleQueryResponse struct {

    // 明细数量
    
    Total   int64 `json:"total,omitempty"`
    

    // 数据主体
    
    Datas  struct {
        AlibabaMjOcBigposBanksaleQueryData  []AlibabaMjOcBigposBanksaleQueryData `json:"alibaba_mj_oc_bigpos_banksale_query_data,omitempty"`
    } `json:"datas,omitempty"`
    

}
*/

type AlibabaMjOcBigposBanksaleQueryResponse struct {

    // 明细数量
    Total   int64 `json:"total,omitempty"`

    // 数据主体
    Datas   []AlibabaMjOcBigposBanksaleQueryData `json:"datas,omitempty"`

}
