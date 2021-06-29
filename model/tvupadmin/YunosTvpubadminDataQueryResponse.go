package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
魔盒统计数据查询接口 APIResponse
yunos.tvpubadmin.data.query

用于华数查询魔盒上的一些用户统计数据
*/
type YunosTvpubadminDataQueryAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDataQueryResponse
}

type YunosTvpubadminDataQueryResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_data_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *DfPageResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
