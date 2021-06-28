package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取修改的推广组ID APIResponse
taobao.simba.adgroupids.changed.get

获取修改的推广组ID
*/
type TaobaoSimbaAdgroupidsChangedGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupidsChangedGetResponse
}

type TaobaoSimbaAdgroupidsChangedGetResponse struct {
    XMLName xml.Name `xml:"simba_adgroupids_changed_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广组ID列表
    
    ChangedAdgroupids   []int64 `json:"changed_adgroupids,omitempty" xml:"changed_adgroupids>int64,omitempty"`
    
    
}
