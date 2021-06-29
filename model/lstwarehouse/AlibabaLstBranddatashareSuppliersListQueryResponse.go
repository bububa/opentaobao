package lstwarehouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌数据授权的供应商列表 APIResponse
alibaba.lst.branddatashare.suppliers.list.query

品牌商查询品牌数据授权的供应商列表
*/
type AlibabaLstBranddatashareSuppliersListQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstBranddatashareSuppliersListQueryResponse
}

type AlibabaLstBranddatashareSuppliersListQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_branddatashare_suppliers_list_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值列表
    
    ContentList   []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
    
    
}
