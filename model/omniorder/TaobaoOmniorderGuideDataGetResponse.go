package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全渠道导购产品数据 API返回值 
taobao.omniorder.guide.data.get

获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
*/
type TaobaoOmniorderGuideDataGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderGuideDataGetResponse
}

// 获取全渠道导购产品数据 成功返回结果
type TaobaoOmniorderGuideDataGetResponse struct {
    XMLName xml.Name `xml:"omniorder_guide_data_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 拉取的数据数组，如果为空，表示数据拉取完毕。拉取的数据字段包括打点时间、商家id、商品id和门店id等，传入的类型不同，返回的字段有所不同，可以根据具体类型的返回结果具体处理
    DataList   []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
}
