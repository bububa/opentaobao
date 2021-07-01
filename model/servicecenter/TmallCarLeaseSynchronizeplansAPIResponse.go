package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步租赁方案 API返回值 
tmall.car.lease.synchronizeplans

租赁公司同步还款计划
*/
type TmallCarLeaseSynchronizeplansAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseSynchronizeplansAPIResponseModel
}

// 同步租赁方案 成功返回结果
type TmallCarLeaseSynchronizeplansAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_car_lease_synchronizeplans_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 如果success为true,表示成功，如果success为false,需要获取msg_code,msg_info,具体的错误码文档提供
    Result   *TmallCarLeaseSynchronizeplansResult `json:"result,omitempty" xml:"result,omitempty"`
}
