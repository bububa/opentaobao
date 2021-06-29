package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级分页查询 API返回值 
yunos.osupdate.osfota.query

分页查询osoupdate系统升级列表
*/
type YunosOsupdateOsfotaQueryAPIResponse struct {
    model.CommonResponse
    YunosOsupdateOsfotaQueryResponse
}

// 系统升级分页查询 成功返回结果
type YunosOsupdateOsfotaQueryResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_osfota_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *YunosOsupdateOsfotaQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
