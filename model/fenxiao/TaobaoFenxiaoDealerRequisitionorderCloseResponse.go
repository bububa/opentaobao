package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商/分销商关闭采购申请/经销采购单 APIResponse
taobao.fenxiao.dealer.requisitionorder.close

供应商或分销商关闭采购申请/经销采购单
*/
type TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_dealer_requisitionorder_close_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作是否成功。true：成功；false：失败。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"