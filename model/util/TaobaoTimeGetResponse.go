package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取淘宝系统当前时间 APIResponse
taobao.time.get

获取淘宝系统当前时间
*/
type TaobaoTimeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTimeGetResponse `json:"time_get_response,omitempty"` 
    TaobaoTimeGetResponse
}

/* model for simplify = false
type TaobaoTimeGetResponse struct {

    // 淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss
    
    Time   string `json:"time,omitempty"`
    

}
*/

type TaobaoTimeGetResponse struct {

    // 淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss
    Time   string `json:"time,omitempty"`

}
