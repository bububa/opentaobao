package dmp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDmpCrowdsGetAPIResponse
查询人群服务 API返回值
taobao.dmp.crowds.get

查询人群服务 */
type TaobaoDmpCrowdsGetAPIResponse struct {
	model.CommonResponse
	TaobaoDmpCrowdsGetAPIResponseModel
}

// TaobaoDmpCrowdsGetAPIResponseModel is 查询人群服务 成功返回结果
type TaobaoDmpCrowdsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"dmp_crowds_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Results []DmpCrowdDto `json:"results,omitempty" xml:"results>dmp_crowd_dto,omitempty"`
}
