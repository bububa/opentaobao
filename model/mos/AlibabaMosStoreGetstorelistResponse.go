package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据屏编号获取专柜集 API返回值 
alibaba.mos.store.getstorelist

根据屏编号获取专柜集
*/
type AlibabaMosStoreGetstorelistAPIResponse struct {
    model.CommonResponse
    AlibabaMosStoreGetstorelistResponse
}

// 根据屏编号获取专柜集 成功返回结果
type AlibabaMosStoreGetstorelistResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_store_getstorelist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaMosStoreGetstorelistResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
