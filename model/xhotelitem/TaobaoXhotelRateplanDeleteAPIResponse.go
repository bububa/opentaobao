package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan删除 API返回值 
taobao.xhotel.rateplan.delete

酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
*/
type TaobaoXhotelRateplanDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateplanDeleteAPIResponseModel
}

// 价格计划rateplan删除 成功返回结果
type TaobaoXhotelRateplanDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_rateplan_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoXhotelRateplanDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
