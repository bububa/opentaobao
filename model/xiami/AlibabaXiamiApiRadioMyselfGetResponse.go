package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
我的电台 APIResponse
alibaba.xiami.api.radio.myself.get

我的电台
*/
type AlibabaXiamiApiRadioMyselfGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_radio_myself_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 歌曲列表
    
    Data   []Song `json:"data,omitempty" xml:"