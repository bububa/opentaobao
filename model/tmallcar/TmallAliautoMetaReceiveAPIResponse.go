package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallAliautoMetaReceiveAPIResponse
汽车说明书元数据上传 API返回值
tmall.aliauto.meta.receive

天猫汽车对外提供的汽车资源元数据上传接口 */
type TmallAliautoMetaReceiveAPIResponse struct {
	model.CommonResponse
	TmallAliautoMetaReceiveAPIResponseModel
}

// TmallAliautoMetaReceiveAPIResponseModel is 汽车说明书元数据上传 成功返回结果
type TmallAliautoMetaReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_meta_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallAliautoMetaReceiveResult `json:"result,omitempty" xml:"result,omitempty"`
}
