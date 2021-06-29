package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取桌面升级任务 APIResponse
yunos.osupdate.appversion.query

分页获取桌面升级任务
*/
type YunosOsupdateAppversionQueryAPIResponse struct {
    model.CommonResponse
    YunosOsupdateAppversionQueryResponse
}

type YunosOsupdateAppversionQueryResponse struct {
    XMLName xml.Name `xml:"yunos_osupdate_appversion_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // resultList
    
    ResultList   []TvAppVersion `json:"result_list,omitempty" xml:"result_list>tv_app_version,omitempty"`
    
    
    // 总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
