package ioti

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下发厂测初始化图片 APIResponse
alibaba.it.esl.eslimage.sendimage

工厂对生产出的电子价签进行全流程功能测试，能将出场图片通过ESL系统初始化到电子价签中
*/
type AlibabaItEslEslimageSendimageAPIResponse struct {
    model.CommonResponse
    AlibabaItEslEslimageSendimageResponse
}

type AlibabaItEslEslimageSendimageResponse struct {
    XMLName xml.Name `xml:"alibaba_it_esl_eslimage_sendimage_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // resultmsg
    
    Resultmsg   string `json:"resultmsg,omitempty" xml:"resultmsg,omitempty"`

    
}
