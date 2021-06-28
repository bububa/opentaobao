package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息接口 APIResponse
alibaba.ele.fengniao.chainstore.update

修改门店的经纬度，文本地址，电话，门店名
*/
type AlibabaEleFengniaoChainstoreUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaEleFengniaoChainstoreUpdateResponse `json:"alibaba_ele_fengniao_chainstore_update_response,omitempty"` 
    AlibabaEleFengniaoChainstoreUpdateResponse
}

/* model for simplify = false
type AlibabaEleFengniaoChainstoreUpdateResponse struct {

    // 出参
    
    Message   string `json:"message,omitempty"`
    

}
*/

type AlibabaEleFengniaoChainstoreUpdateResponse struct {

    // 出参
    Message   string `json:"message,omitempty"`

}
