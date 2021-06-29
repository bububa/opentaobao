package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询拣货单中的药品信息 API返回值 
wdk.wms.pick.medicine.query

联营商药机查询拣货单中的药品信息
*/
type WdkWmsPickMedicineQueryAPIResponse struct {
    model.CommonResponse
    WdkWmsPickMedicineQueryResponse
}

// 查询拣货单中的药品信息 成功返回结果
type WdkWmsPickMedicineQueryResponse struct {
    XMLName xml.Name `xml:"wdk_wms_pick_medicine_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *WdkWmsPickMedicineQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
