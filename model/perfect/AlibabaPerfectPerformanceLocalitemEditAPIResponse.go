package perfect

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaperfectperformancelocalitemeditAPIResponse 同城购定制发品编辑 API返回值
// alibaba.perfect.performance.localitem.edit
//
// 同城购业务定制化发品接口，同城购业务线专用
type AlibabaperfectperformancelocalitemeditAPIResponse struct {
	model.CommonResponse
	AlibabaperfectperformancelocalitemeditAPIResponseModel
}

// AlibabaperfectperformancelocalitemeditAPIResponseModel is 同城购定制发品编辑 成功返回结果
type AlibabaperfectperformancelocalitemeditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_perfect_performance_localitem_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
