package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系新增 API返回值 
taobao.place.storerelatesub.add

门店和子门店关系新增
*/
type TaobaoPlaceStorerelatesubAddAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStorerelatesubAddResponse
}

// 门店和子门店关系新增 成功返回结果
type TaobaoPlaceStorerelatesubAddResponse struct {
    XMLName xml.Name `xml:"place_storerelatesub_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
