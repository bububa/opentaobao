package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterspserviceorderepocreceiveAPIResponse 电子保单数据接口 API返回值
// tmall.servicecenter.spserviceorder.epoc.receive
//
// 电子保单数据回传接口
type TmallservicecenterspserviceorderepocreceiveAPIResponse struct {
	model.CommonResponse
	TmallservicecenterspserviceorderepocreceiveAPIResponseModel
}

// TmallservicecenterspserviceorderepocreceiveAPIResponseModel is 电子保单数据接口 成功返回结果
type TmallservicecenterspserviceorderepocreceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_epoc_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}
