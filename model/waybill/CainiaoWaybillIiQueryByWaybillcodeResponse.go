package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过面单号查询面单信息 APIResponse
cainiao.waybill.ii.query.by.waybillcode

通过面单号查看面单号的当前状态，如签收、发货、失效等。
*/
type CainiaoWaybillIiQueryByWaybillcodeAPIResponse struct {
    model.CommonResponse
    CainiaoWaybillIiQueryByWaybillcodeResponse
}

type CainiaoWaybillIiQueryByWaybillcodeResponse struct {
    XMLName xml.Name `xml:"cainiao_waybill_ii_query_by_waybillcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询返回值
    
    Modules   []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty" xml:"modules>waybill_cloud_print_with_result_desc_response,omitempty"`
    
    
}
