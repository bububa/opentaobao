package ihome

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoihomeadvancepicuploadAPIResponse ihome图片上传 API返回值
// taobao.ihome.advancepic.upload
//
// ihome 定制业务编辑器投稿素材上传
type TaobaoihomeadvancepicuploadAPIResponse struct {
	model.CommonResponse
	TaobaoihomeadvancepicuploadAPIResponseModel
}

// TaobaoihomeadvancepicuploadAPIResponseModel is ihome图片上传 成功返回结果
type TaobaoihomeadvancepicuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"ihome_advancepic_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *TaobaoihomeadvancepicuploadApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
