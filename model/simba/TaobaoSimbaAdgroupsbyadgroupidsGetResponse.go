package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广组 API返回值 
taobao.simba.adgroupsbyadgroupids.get

批量得到推广组
*/
type TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupsbyadgroupidsGetResponse
}

// 批量得到推广组 成功返回结果
type TaobaoSimbaAdgroupsbyadgroupidsGetResponse struct {
    XMLName xml.Name `xml:"simba_adgroupsbyadgroupids_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回的推广组分页对象
    Adgroups   *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}
