package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcEserviceAppointListAPIResponse
交互卡片预约信息读取接口 API返回值
taobao.oc.eservice.appoint.list

允许外部的isv通过该接口读取门店预约信息 */
type TaobaoOcEserviceAppointListAPIResponse struct {
	model.CommonResponse
	TaobaoOcEserviceAppointListAPIResponseModel
}

// TaobaoOcEserviceAppointListAPIResponseModel is 交互卡片预约信息读取接口 成功返回结果
type TaobaoOcEserviceAppointListAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_eservice_appoint_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的预约信息，多个预约信息按照预约时间升序排序
	Models []O2oAppointInfoDto `json:"models,omitempty" xml:"models>o2o_appoint_info_dto,omitempty"`
	// 返回的预约信息数目
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}
