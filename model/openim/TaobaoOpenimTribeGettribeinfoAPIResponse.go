package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取群信息 API返回值 
taobao.openim.tribe.gettribeinfo

获取群信息
*/
type TaobaoOpenimTribeGettribeinfoAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribeGettribeinfoAPIResponseModel
}

// 获取群信息 成功返回结果
type TaobaoOpenimTribeGettribeinfoAPIResponseModel struct {
    XMLName xml.Name `xml:"openim_tribe_gettribeinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 群信息
    TribeInfo   *TribeInfo `json:"tribe_info,omitempty" xml:"tribe_info,omitempty"`
}
