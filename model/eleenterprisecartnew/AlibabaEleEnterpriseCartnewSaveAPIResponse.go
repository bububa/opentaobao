package eleenterprisecartnew

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新版创建购物车 API返回值 
alibaba.ele.enterprise.cartnew.save

新版创建购物车
*/
type AlibabaEleEnterpriseCartnewSaveAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseCartnewSaveAPIResponseModel
}

// 新版创建购物车 成功返回结果
type AlibabaEleEnterpriseCartnewSaveAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_cartnew_save_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 状态码
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
    // 状态消息
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
    // 系统自动生成
    EnterpriseDatas   *EnterpriseData `json:"enterprise_datas,omitempty" xml:"enterprise_datas,omitempty"`
    // 请求id
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}
