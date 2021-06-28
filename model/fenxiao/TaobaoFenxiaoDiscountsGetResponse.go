package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取折扣信息 APIResponse
taobao.fenxiao.discounts.get

查询折扣信息
*/
type TaobaoFenxiaoDiscountsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_discounts_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 折扣数据结构
    
    Discounts   []Discount `json:"discounts,omitempty" xml:"