package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步天猫配送人员信息 APIResponse
tmall.nr.notice.goods.ready

接收商家的配送人员信息，和第三公司信息及提货码
*/
type TmallNrNoticeGoodsReadyAPIResponse struct {
    model.CommonResponse
    TmallNrNoticeGoodsReadyResponse
}

type TmallNrNoticeGoodsReadyResponse struct {
    XMLName xml.Name `xml:"tmall_nr_notice_goods_ready_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // resultData
    
    ResultData   bool `json:"result_data,omitempty" xml:"result_data,omitempty"`

    
    // errorMessage
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`

    
    // errorCode
    
    ErrorCode2   string `json:"error_code2,omitempty" xml:"error_code2,omitempty"`

    
    // isSuccess
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
