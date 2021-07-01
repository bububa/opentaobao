package eleenterpriseordernew

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态查询接口 API返回值 
alibaba.ele.enterprise.ordernew.getstatus

订单状态查询接口
*/
type AlibabaEleEnterpriseOrdernewGetstatusAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseOrdernewGetstatusAPIResponseModel
}

// 订单状态查询接口 成功返回结果
type AlibabaEleEnterpriseOrdernewGetstatusAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_getstatus_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值信息
    EnterpriseData   *StandardOrderTrackingInfoDto `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
    // 响应code
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
    // 响应信息
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
    // 请求id
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}
