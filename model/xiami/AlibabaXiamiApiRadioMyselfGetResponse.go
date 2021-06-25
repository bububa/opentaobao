package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
我的电台 APIResponse
alibaba.xiami.api.radio.myself.get

我的电台
*/
type AlibabaXiamiApiRadioMyselfGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiRadioMyselfGetResponse `json:"alibaba_xiami_api_radio_myself_get_response,omitempty"`
}

type AlibabaXiamiApiRadioMyselfGetResponse struct {

    // 歌曲列表
    Data   []Song `json:"data,omitempty"`

}
