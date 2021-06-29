package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票单接口batchPushTicket API请求
alibaba.damai.mev.open.batchpushticket

批量推送票单
*/
type AlibabaDamaiMevOpenBatchpushticketRequest struct {
    model.Params
    // 入参thirdTicketSetOpenParamList
    _thirdTicketSetOpenParamList   []ThirdTicketPushOpenParam
}

// 初始化AlibabaDamaiMevOpenBatchpushticketRequest对象
func NewAlibabaDamaiMevOpenBatchpushticketRequest() *AlibabaDamaiMevOpenBatchpushticketRequest{
    return &AlibabaDamaiMevOpenBatchpushticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenBatchpushticketRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.batchpushticket"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenBatchpushticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ThirdTicketSetOpenParamList Setter
// 入参thirdTicketSetOpenParamList
func (r *AlibabaDamaiMevOpenBatchpushticketRequest) SetThirdTicketSetOpenParamList(_thirdTicketSetOpenParamList []ThirdTicketPushOpenParam) error {
    r._thirdTicketSetOpenParamList = _thirdTicketSetOpenParamList
    r.Set("third_ticket_set_open_param_list", _thirdTicketSetOpenParamList)
    return nil
}

// ThirdTicketSetOpenParamList Getter
func (r AlibabaDamaiMevOpenBatchpushticketRequest) GetThirdTicketSetOpenParamList() []ThirdTicketPushOpenParam {
    return r._thirdTicketSetOpenParamList
}
