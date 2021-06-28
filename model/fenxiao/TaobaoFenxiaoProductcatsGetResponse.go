package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询产品线列表 APIResponse
taobao.fenxiao.productcats.get

查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
*/
type TaobaoFenxiaoProductcatsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_productcats_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果记录数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"