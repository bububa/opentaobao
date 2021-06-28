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
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_distributor_items_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果记录数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"