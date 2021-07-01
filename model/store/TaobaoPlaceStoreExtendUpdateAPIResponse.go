package store

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店拓展信息更新接口 API返回值 
taobao.place.store.extend.update

更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
*/
type TaobaoPlaceStoreExtendUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPlaceStoreExtendUpdateAPIResponseModel
}

// 商户门店拓展信息更新接口 成功返回结果
type TaobaoPlaceStoreExtendUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"place_store_extend_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoPlaceStoreExtendUpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
