package logistic

import (
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
    Response *TaobaoAreasGetResponse `json:"taobao_areas_get_response,omitempty"`
}

type TaobaoAreasGetResponse struct {

    // 地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息 。
    Areas   []Area `json:"areas,omitempty"`

}
