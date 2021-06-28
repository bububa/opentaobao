package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品下载记录 APIResponse
taobao.fenxiao.distributor.items.get

供应商查询分销商商品下载记录。
*/
type TaobaoFenxiaoDistributorItemsGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoDistributorItemsGetResponse
}

type TaobaoFenxiaoDistributorItemsGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_distributor_items_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果记录数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 下载记录对象
    
    Records   []FenxiaoItemRecord `json:"records,omitempty" xml:"records>fenxiao_item_record,omitempty"`
    
    
}
