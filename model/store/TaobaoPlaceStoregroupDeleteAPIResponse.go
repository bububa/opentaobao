package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除门店库 API返回值 
taobao.place.storegroup.delete

删除门店库
*/
type TaobaoPlaceStoregroupDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoregroupDeleteAPIResponseModel
}

// 删除门店库 成功返回结果
type TaobaoPlaceStoregroupDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"place_storegroup_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TopBatchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
