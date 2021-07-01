package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川登录二次验证 API返回值 
taobao.baichuan.openaccount.logindoublecheck

百川登录二次验证
*/
type TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountLogindoublecheckAPIResponseModel
}

// 百川登录二次验证 成功返回结果
type TaobaoBaichuanOpenaccountLogindoublecheckAPIResponseModel struct {
    XMLName xml.Name `xml:"baichuan_openaccount_logindoublecheck_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
