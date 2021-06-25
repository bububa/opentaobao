package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询解析淘口令 APIResponse
alibaba.baichuan.taopassword.query

查询，解析淘口令
*/
type AlibabaBaichuanTaopasswordQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaBaichuanTaopasswordQueryResponse `json:"alibaba_baichuan_taopassword_query_response,omitempty"`
}

type AlibabaBaichuanTaopasswordQueryResponse struct {

    // result
    Result   *BcTaoPasswordResult `json:"result,omitempty"`

}
