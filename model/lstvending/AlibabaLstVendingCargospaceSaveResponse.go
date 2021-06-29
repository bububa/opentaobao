package lstvending

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机货道数据回流 API返回值 
alibaba.lst.vending.cargospace.save

自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
*/
type AlibabaLstVendingCargospaceSaveAPIResponse struct {
    model.CommonResponse
    AlibabaLstVendingCargospaceSaveResponse
}

// 自动售卖机货道数据回流 成功返回结果
type AlibabaLstVendingCargospaceSaveResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_vending_cargospace_save_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果集
    Result   *MultiResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
