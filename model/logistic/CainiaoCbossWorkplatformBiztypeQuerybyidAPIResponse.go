package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaocbossworkplatformbiztypequerybyidAPIResponse 菜鸟工单平台根据业务类型id查询业务类型详细信息 API返回值
// cainiao.cboss.workplatform.biztype.querybyid
//
// 菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
type CainiaocbossworkplatformbiztypequerybyidAPIResponse struct {
	model.CommonResponse
	CainiaocbossworkplatformbiztypequerybyidAPIResponseModel
}

// CainiaocbossworkplatformbiztypequerybyidAPIResponseModel is 菜鸟工单平台根据业务类型id查询业务类型详细信息 成功返回结果
type CainiaocbossworkplatformbiztypequerybyidAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cboss_workplatform_biztype_querybyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CainiaocbossworkplatformbiztypequerybyidResult `json:"result,omitempty" xml:"result,omitempty"`
}
