package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单号查询电子面单通接口 APIResponse
cainiao.waybill.ii.query.by.tradecode

通过订单号查看面单的信息
*/
type CainiaoWaybillIiQueryByTradecodeAPIResponse struct {
    model.CommonResponse
    CainiaoWaybillIiQueryByTradecodeResponse
}

type CainiaoWaybillIiQueryByTradecodeResponse struct {
    XMLName xml.Name `xml:"cainiao_waybill_ii_query_by_tradecode_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询返回值
    
    Modules   []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_with_result_desc_response,omitempty"`
    
    
}
