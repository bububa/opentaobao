package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取云货架版本信息 APIResponse
alibaba.mos.store.getcloudshelfversion

根据屏编号获取云货架版本信息
*/
type AlibabaMosStoreGetcloudshelfversionAPIResponse struct {
    model.CommonResponse
    AlibabaMosStoreGetcloudshelfversionResponse
}

type AlibabaMosStoreGetcloudshelfversionResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_store_getcloudshelfversion_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaMosStoreGetcloudshelfversionResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
