package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取云货架版本信息 APIResponse
alibaba.mos.store.getcloudshelfversion

根据屏编号获取云货架版本信息
*/
type AlibabaMosStoreGetcloudshelfversionAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosStoreGetcloudshelfversionResponse `json:"alibaba_mos_store_getcloudshelfversion_response,omitempty"`
}

type AlibabaMosStoreGetcloudshelfversionResponse struct {

    // result
    Result   *AlibabaMosStoreGetcloudshelfversionResultDo `json:"result,omitempty"`

}
