package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户常用发票信息查询接口 APIResponse
taobao.xhotel.commoninvoice.list.vtwo

获取用户常用发票信息接口
*/
type TaobaoXhotelCommoninvoiceListVtwoAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelCommoninvoiceListVtwoResponse
}

type TaobaoXhotelCommoninvoiceListVtwoResponse struct {
    XMLName xml.Name `xml:"xhotel_commoninvoice_list_vtwo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果集
    
    Result   *TaobaoXhotelCommoninvoiceListVtwoResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
