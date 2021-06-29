package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店库修改基本信息 API返回值 
taobao.place.storegroup.update

门店库修改基本信息
*/
type TaobaoPlaceStoregroupUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoregroupUpdateResponse
}

// 门店库修改基本信息 成功返回结果
type TaobaoPlaceStoregroupUpdateResponse struct {
    XMLName xml.Name `xml:"place_storegroup_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TopBatchResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
