package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取等级列表 APIResponse
alibaba.westcrm.grade.get

获取会员卡等级列表
*/
type AlibabaWestcrmGradeGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmGradeGetResponse
}

type AlibabaWestcrmGradeGetResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_grade_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
