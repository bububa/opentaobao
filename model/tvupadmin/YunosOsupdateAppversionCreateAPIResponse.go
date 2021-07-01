package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建应用升级任务 API返回值 
yunos.osupdate.appversion.create

创建应用升级任务
*/
type YunosOsupdateAppversionCreateAPIResponse struct {
    model.CommonResponse
    YunosOsupdateAppversionCreateAPIResponseModel
}

// 创建应用升级任务 成功返回结果
type YunosOsupdateAppversionCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"yunos_osupdate_appversion_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
