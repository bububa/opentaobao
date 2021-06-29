package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据推广单元id获取关键词 APIResponse
taobao.simba.keyword.findbyadgroupid

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyadgroupidAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordFindbyadgroupidResponse
}

type TaobaoSimbaKeywordFindbyadgroupidResponse struct {
    XMLName xml.Name `xml:"simba_keyword_findbyadgroupid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 整体的返回值
    
    Results   []SiriusBidwordDto `json:"results,omitempty" xml:"results>sirius_bidword_dto,omitempty"`
    
    
    // 错误原因
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
