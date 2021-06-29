package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部小件员休息 APIResponse
cainiao.nborderfront.user.outside.queryoutsideuser

采用SPI方式查询外部公司的小件员信息
*/
type CainiaoNborderfrontUserOutsideQueryoutsideuserAPIResponse struct {
    model.CommonResponse
    CainiaoNborderfrontUserOutsideQueryoutsideuserResponse
}

type CainiaoNborderfrontUserOutsideQueryoutsideuserResponse struct {
    XMLName xml.Name `xml:"cainiao_nborderfront_user_outside_queryoutsideuser_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回错误编码
    
    RespCode   string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`

    
    // userInfo
    
    UserInfo   *CainiaoNborderfrontUserOutsideQueryoutsideuserStruct `json:"user_info,omitempty" xml:"user_info,omitempty"`

    
}
