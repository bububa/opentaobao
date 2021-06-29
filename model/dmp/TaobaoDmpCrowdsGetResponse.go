package dmp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询人群服务 API返回值 
taobao.dmp.crowds.get

查询人群服务
*/
type TaobaoDmpCrowdsGetAPIResponse struct {
    model.CommonResponse
    TaobaoDmpCrowdsGetResponse
}

// 查询人群服务 成功返回结果
type TaobaoDmpCrowdsGetResponse struct {
    XMLName xml.Name `xml:"dmp_crowds_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 1
    Results   []DmpCrowdDTO `json:"results,omitempty" xml:"results>dmp_crowd_dto,omitempty"`
}
