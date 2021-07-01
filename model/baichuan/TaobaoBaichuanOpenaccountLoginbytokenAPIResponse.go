package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
百川TOKEN 登录 API返回值 
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录
*/
type TaobaoBaichuanOpenaccountLoginbytokenAPIResponse struct {
    model.CommonResponse
    TaobaoBaichuanOpenaccountLoginbytokenAPIResponseModel
}

// 百川TOKEN 登录 成功返回结果
type TaobaoBaichuanOpenaccountLoginbytokenAPIResponseModel struct {
    XMLName xml.Name `xml:"baichuan_openaccount_loginbytoken_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // name
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
