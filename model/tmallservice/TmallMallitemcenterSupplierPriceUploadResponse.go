package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务商服务报价上传 APIResponse
tmall.mallitemcenter.supplier.price.upload

天猫服务商上传服务价格
*/
type TmallMallitemcenterSupplierPriceUploadAPIResponse struct {
    model.CommonResponse
    TmallMallitemcenterSupplierPriceUploadResponse
}

type TmallMallitemcenterSupplierPriceUploadResponse struct {
    XMLName xml.Name `xml:"tmall_mallitemcenter_supplier_price_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TmallMallitemcenterSupplierPriceUploadResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
