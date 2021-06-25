package ma

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建二维码/短连接 APIResponse
taobao.wireless.xcode.create

创建码平台的普通二维码或者长连接转短连接服务
*/
type TaobaoWirelessXcodeCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWirelessXcodeCreateResponse `json:"taobao_wireless_xcode_create_response,omitempty"`
}

type TaobaoWirelessXcodeCreateResponse struct {

    // 创建二维码/短连接 返回信息
    Xcode   *XCodeTo `json:"xcode,omitempty"`

}
