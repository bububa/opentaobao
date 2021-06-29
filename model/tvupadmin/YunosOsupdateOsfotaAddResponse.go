package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加系统升级任务 API返回值 
yunos.osupdate.osfota.add

添加osupdate系统升级任务
*/
type YunosOsupdateOsfotaAddAPIResponse struct {
    model.CommonResponse
    YunosOsupdateOsfotaAddResponse
}

// 添加系统升级任务 成功返回结果
type YunosOsupdateOsfotaAddResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_osfota_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *YunosOsupdateOsfotaAddResult `json:"result,omitempty" xml:"result,omitempty"`
}
