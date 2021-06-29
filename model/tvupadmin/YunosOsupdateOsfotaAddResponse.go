package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加系统升级任务 APIResponse
yunos.osupdate.osfota.add

添加osupdate系统升级任务
*/
type YunosOsupdateOsfotaAddAPIResponse struct {
    model.CommonResponse
    YunosOsupdateOsfotaAddResponse
}

type YunosOsupdateOsfotaAddResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_osfota_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *YunosOsupdateOsfotaAddResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
