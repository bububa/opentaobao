package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存网点 APIResponse
alibaba.ssc.supplyplatform.servicestore.save

网点创建、修改
*/
type AlibabaSscSupplyplatformServicestoreSaveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ssc_supplyplatform_servicestore_save_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaSscSupplyplatformServicestoreSaveResult `json:"result,omitempty" xml:"