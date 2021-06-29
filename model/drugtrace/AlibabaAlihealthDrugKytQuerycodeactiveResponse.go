package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询码是否激活 API返回值 
alibaba.alihealth.drug.kyt.querycodeactive

查询码是否激活
*/
type AlibabaAlihealthDrugKytQuerycodeactiveAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytQuerycodeactiveResponse
}

// 查询码是否激活 成功返回结果
type AlibabaAlihealthDrugKytQuerycodeactiveResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_querycodeactive_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 未激活的码
    Models   []string `json:"models,omitempty" xml:"models>string,omitempty"`
    // 错误码(BILL_DECODE_ERROR 单据转码失败 2.BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 是否成功(true 成功 ,false失败)
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
