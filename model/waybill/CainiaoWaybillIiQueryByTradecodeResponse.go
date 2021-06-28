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
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_waybill_ii_query_by_tradecode_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询返回值
    
    Modules   []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty" xml:"