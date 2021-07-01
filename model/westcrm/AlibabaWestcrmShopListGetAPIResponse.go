package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商圈商户信息列表 API返回值 
alibaba.westcrm.shop.list.get

获取商圈商户信息列表
*/
type AlibabaWestcrmShopListGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmShopListGetAPIResponseModel
}

// 获取商圈商户信息列表 成功返回结果
type AlibabaWestcrmShopListGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_westcrm_shop_list_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象封装
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
