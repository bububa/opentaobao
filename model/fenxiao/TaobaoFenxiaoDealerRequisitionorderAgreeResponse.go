package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商/分销商通过采购申请/经销采购单申请 APIResponse
taobao.fenxiao.dealer.requisitionorder.agree

供应商或分销商通过采购申请/经销采购单审核
*/
type TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoDealerRequisitionorderAgreeResponse
}

type TaobaoFenxiaoDealerRequisitionorderAgreeResponse struct {
    XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_agree_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作是否成功。true：成功；false：失败。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
