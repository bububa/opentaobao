package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取加密公钥 APIResponse
alibaba.alihealth.drugcode.drugfactory.getencrptypk

获取服务端给药厂用来加密的公钥
*/
type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkResponse
}

type AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_getencrptypk_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 公钥证书
    
    Model   string `json:"model,omitempty" xml:"model,omitempty"`

    
    // 操作码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 操作说明
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
