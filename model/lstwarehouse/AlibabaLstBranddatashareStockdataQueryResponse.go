package lstwarehouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌商品实仓库存/周转效能 APIResponse
alibaba.lst.branddatashare.stockdata.query

品牌商查询授权供应商实仓库存数据
*/
type AlibabaLstBranddatashareStockdataQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstBranddatashareStockdataQueryResponse
}

type AlibabaLstBranddatashareStockdataQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_branddatashare_stockdata_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstBranddatashareStockdataQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
