package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口batchPushTicket APIRequest
alibaba.damai.mev.open.batchpushticket

批量推送票单
*/
type AlibabaDamaiMevOpenBatchpushticketRequest struct {
    model.Params

    // 入参thirdTicketSetOpenParamList
    thirdTicketSetOpenParamList   []ThirdTicketPushOpenParam 

}

func NewAlibabaDamaiMevOpenBatchpushticketRequest() *AlibabaDamaiMevOpenBatchpushticketRequest{
    return &AlibabaDamaiMevOpenBatchpushticketRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenBatchpushticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.batchpushticket"
}

func (r AlibabaDamaiMevOpenBatchpushticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenBatchpushticketRequest) SetThirdTicketSetOpenParamList(thirdTicketSetOpenParamList []ThirdTicketPushOpenParam) error {
    r.thirdTicketSetOpenParamList = thirdTicketSetOpenParamList
    r.Set("third_ticket_set_open_param_list", thirdTicketSetOpenParamList)
    return nil
}

func (r AlibabaDamaiMevOpenBatchpushticketRequest) GetThirdTicketSetOpenParamList() []ThirdTicketPushOpenParam {
    return r.thirdTicketSetOpenParamList
}

