package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询地址区域 APIResponse
taobao.areas.get

查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
<a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/" target="_blank"> http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/</a>
*/
type TaobaoAreasGetAPIResponse struct {
    model.CommonResponse
    TaobaoAreasGetResponse
}

type TaobaoAreasGetResponse struct {
    XMLName xml.Name `xml:"areas_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。
    
    Areas   []Area `json:"areas,omitempty" xml:"areas>area,omitempty"`
    
    
}
