package alisports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponse 阿里体育运动数据同步统一接口 API返回值
// alibaba.alisports.datacenter.datasync.sportsdatas
//
// 给单方提供同步运动数据到阿里体育的接口
type AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponseModel
}

// AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponseModel is 阿里体育运动数据同步统一接口 成功返回结果
type AlibabaAlisportsDatacenterDatasyncSportsdatasAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_datacenter_datasync_sportsdatas_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	RsMsg string `json:"rs_msg,omitempty" xml:"rs_msg,omitempty"`
	// 返回码
	RsCode string `json:"rs_code,omitempty" xml:"rs_code,omitempty"`
	// 是否成功
	RsSuccess bool `json:"rs_success,omitempty" xml:"rs_success,omitempty"`
	// 返回值
	RsModel bool `json:"rs_model,omitempty" xml:"rs_model,omitempty"`
}
