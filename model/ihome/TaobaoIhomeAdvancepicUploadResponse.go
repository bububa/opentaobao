package ihome

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ihome图片上传 APIResponse
taobao.ihome.advancepic.upload

ihome 定制业务编辑器投稿素材上传
*/
type TaobaoIhomeAdvancepicUploadAPIResponse struct {
    model.CommonResponse
    TaobaoIhomeAdvancepicUploadResponse
}

type TaobaoIhomeAdvancepicUploadResponse struct {
    XMLName xml.Name `xml:"ihome_advancepic_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 根据站点名称查询产品
    
    Result   *TaobaoIhomeAdvancepicUploadApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
