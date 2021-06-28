package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商同步接口 APIResponse
taobao.qimen.supplier.synchronize

这个接口用来同步供应商信息
*/
type TaobaoQimenSupplierSynchronizeAPIResponse struct {
    model.CommonResponse
    TaobaoQimenSupplierSynchronizeResponse
}

type TaobaoQimenSupplierSynchronizeResponse struct {
    XMLName xml.Name `xml:"qimen_supplier_synchronize_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
