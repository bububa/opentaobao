package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元增加搜索人群 APIResponse
taobao.simba.searchcrowd.batch.add

推广单元新增搜索人群
*/
type TaobaoSimbaSearchcrowdBatchAddAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSearchcrowdBatchAddResponse
}

type TaobaoSimbaSearchcrowdBatchAddResponse struct {
    XMLName xml.Name `xml:"simba_searchcrowd_batch_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 定向信息
    
    Adgrouptargetingtags   []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty" xml:"adgrouptargetingtags>adgroup_targeting_tag_dto,omitempty"`
    
    
}
