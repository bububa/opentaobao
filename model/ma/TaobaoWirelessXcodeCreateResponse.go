package ma

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建二维码/短连接 APIResponse
taobao.wireless.xcode.create

创建码平台的普通二维码或者长连接转短连接服务
*/
type TaobaoWirelessXcodeCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wireless_xcode_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建二维码/短连接 返回信息
    
    Xcode   *XCodeTo `json:"xcode,omitempty" xml:"