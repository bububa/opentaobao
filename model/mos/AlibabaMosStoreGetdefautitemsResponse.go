package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取默认状态下商品列表 APIResponse
alibaba.mos.store.getdefautitems

获取默认状态下商品列表
*/
type AlibabaMosStoreGetdefautitemsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_store_getdefautitems_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaMosStoreGetdefautitemsResultDo `json:"result,omitempty" xml:"