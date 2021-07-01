package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询实体对应的服务时间数据 API返回值 
taobao.xhotel.servicetime.get

通过实体来获取对应的服务时间数据
*/
type TaobaoXhotelServicetimeGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelServicetimeGetAPIResponseModel
}

// 查询实体对应的服务时间数据 成功返回结果
type TaobaoXhotelServicetimeGetAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_servicetime_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoXhotelServicetimeGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
