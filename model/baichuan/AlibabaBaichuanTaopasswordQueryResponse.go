package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询解析淘口令 APIResponse
alibaba.baichuan.taopassword.query

查询，解析淘口令
*/
type AlibabaBaichuanTaopasswordQueryAPIResponse struct {
    model.CommonResponse
    AlibabaBaichuanTaopasswordQueryResponse
}

type AlibabaBaichuanTaopasswordQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_baichuan_taopassword_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *BcTaoPasswordResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
