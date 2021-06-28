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
    // Response *AlibabaXiamiApiRadioMyselfGetResponse `json:"alibaba_xiami_api_radio_myself_get_response,omitempty"` 
    AlibabaXiamiApiRadioMyselfGetResponse
}

/* model for simplify = false
type AlibabaXiamiApiRadioMyselfGetResponse struct {

    // 歌曲列表
    
    Data  struct {
        Song  []Song `json:"song,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaXiamiApiRadioMyselfGetResponse struct {

    // 歌曲列表
    Data   []Song `json:"data,omitempty"`

}
