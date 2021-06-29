package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单品搜索人群批量取消溢价 APIResponse
taobao.simba.serchcrowd.batch.delete

删除单品搜索人群溢价功能
*/
type TaobaoSimbaSerchcrowdBatchDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSerchcrowdBatchDeleteResponse
}

type TaobaoSimbaSerchcrowdBatchDeleteResponse struct {
    XMLName xml.Name `xml:"simba_serchcrowd_batch_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    DeleteList   []string `json:"delete_list,omitempty" xml:"delete_list>string,omitempty"`
    
    
    // errorDTOList
    
    ErrorDTOList   []string `json:"error_d_t_o_list,omitempty" xml:"error_d_t_o_list>string,omitempty"`
    
    
    // key
    
    Key   string `json:"key,omitempty" xml:"key,omitempty"`

    
}
