package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautometareceiveAPIResponse 汽车说明书元数据上传 API返回值
// tmall.aliauto.meta.receive
//
// 天猫汽车对外提供的汽车资源元数据上传接口
type TmallaliautometareceiveAPIResponse struct {
	model.CommonResponse
	TmallaliautometareceiveAPIResponseModel
}

// TmallaliautometareceiveAPIResponseModel is 汽车说明书元数据上传 成功返回结果
type TmallaliautometareceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_meta_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallaliautometareceiveResult `json:"result,omitempty" xml:"result,omitempty"`
}
