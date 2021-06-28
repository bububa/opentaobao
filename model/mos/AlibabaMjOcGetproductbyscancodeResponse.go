package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
POS商品查询接口 APIResponse
alibaba.mj.oc.getproductbyscancode

此API用于在银泰商场中，POS端扫码获取商品信息
*/
type AlibabaMjOcGetproductbyscancodeAPIResponse struct {
    model.CommonResponse
    AlibabaMjOcGetproductbyscancodeResponse
}

type AlibabaMjOcGetproductbyscancodeResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_oc_getproductbyscancode_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // data
    
    ProductList   []ScanProduct `json:"product_list,omitempty" xml:"product_list>scan_product,omitempty"`
    
    
}
