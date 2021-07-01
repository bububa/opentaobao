package lstvending

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机商品回传 API返回值 
alibaba.lst.vending.goods.save

零售通自动售卖机商品数据回流。
*/
type AlibabaLstVendingGoodsSaveAPIResponse struct {
    model.CommonResponse
    AlibabaLstVendingGoodsSaveResponse
}

// 自动售卖机商品回传 成功返回结果
type AlibabaLstVendingGoodsSaveResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_vending_goods_save_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果集
    Result   *MultiResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
