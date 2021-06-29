package aliospay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
支付记录批量查询接口 APIResponse
aliyun.alios.pay.record.list

商户用来对账的接口
*/
type AliyunAliosPayRecordListAPIResponse struct {
    model.CommonResponse
    AliyunAliosPayRecordListResponse
}

type AliyunAliosPayRecordListResponse struct {
    XMLName xml.Name `xml:"aliyun_alios_pay_record_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应参数
    
    AliospayResponse   *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`

    
}
