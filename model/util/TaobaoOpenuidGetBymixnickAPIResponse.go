package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过mixnick转换openuid API返回值 
taobao.openuid.get.bymixnick

通过mixnick转换openuid
*/
type TaobaoOpenuidGetBymixnickAPIResponse struct {
    model.CommonResponse
    TaobaoOpenuidGetBymixnickAPIResponseModel
}

// 通过mixnick转换openuid 成功返回结果
type TaobaoOpenuidGetBymixnickAPIResponseModel struct {
    XMLName xml.Name `xml:"openuid_get_bymixnick_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // OpenUID
    OpenUid   string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}
