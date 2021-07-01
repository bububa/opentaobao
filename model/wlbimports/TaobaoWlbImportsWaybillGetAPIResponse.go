package wlbimports

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取运单信息 API返回值 
taobao.wlb.imports.waybill.get

一般进口商家，获取订单的电子面单链接地址
*/
type TaobaoWlbImportsWaybillGetAPIResponse struct {
    model.CommonResponse
    TaobaoWlbImportsWaybillGetAPIResponseModel
}

// 获取运单信息 成功返回结果
type TaobaoWlbImportsWaybillGetAPIResponseModel struct {
    XMLName xml.Name `xml:"wlb_imports_waybill_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 电子面单链接地址
    Waybillurl   string `json:"waybillurl,omitempty" xml:"waybillurl,omitempty"`
}
