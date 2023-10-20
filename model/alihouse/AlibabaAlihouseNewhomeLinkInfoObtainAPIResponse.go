package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse 落地页获取 API返回值
// alibaba.alihouse.newhome.link.info.obtain
//
// 落地页获取
type AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel
}

// AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel is 落地页获取 成功返回结果
type AlibabaAlihouseNewhomeLinkInfoObtainAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_link_info_obtain_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeLinkInfoObtainResult `json:"result,omitempty" xml:"result,omitempty"`
}
