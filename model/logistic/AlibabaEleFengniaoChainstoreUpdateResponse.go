package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息接口 API返回值 
alibaba.ele.fengniao.chainstore.update

修改门店的经纬度，文本地址，电话，门店名
*/
type AlibabaEleFengniaoChainstoreUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoChainstoreUpdateResponse
}

// 修改门店信息接口 成功返回结果
type AlibabaEleFengniaoChainstoreUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_chainstore_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 出参
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
