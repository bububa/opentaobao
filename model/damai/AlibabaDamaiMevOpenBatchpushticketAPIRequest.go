package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenBatchpushticketAPIRequest 大麦换验平台-第三方对外开放-票单接口batchPushTicket API请求
// alibaba.damai.mev.open.batchpushticket
//
// 批量推送票单
type AlibabaDamaiMevOpenBatchpushticketAPIRequest struct {
	model.Params
	// 入参thirdTicketSetOpenParamList
	_thirdTicketSetOpenParamList []ThirdTicketPushOpenParam
}

// NewAlibabaDamaiMevOpenBatchpushticketRequest 初始化AlibabaDamaiMevOpenBatchpushticketAPIRequest对象
func NewAlibabaDamaiMevOpenBatchpushticketRequest() *AlibabaDamaiMevOpenBatchpushticketAPIRequest {
	return &AlibabaDamaiMevOpenBatchpushticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenBatchpushticketAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.batchpushticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenBatchpushticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenBatchpushticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThirdTicketSetOpenParamList is ThirdTicketSetOpenParamList Setter
// 入参thirdTicketSetOpenParamList
func (r *AlibabaDamaiMevOpenBatchpushticketAPIRequest) SetThirdTicketSetOpenParamList(_thirdTicketSetOpenParamList []ThirdTicketPushOpenParam) error {
	r._thirdTicketSetOpenParamList = _thirdTicketSetOpenParamList
	r.Set("third_ticket_set_open_param_list", _thirdTicketSetOpenParamList)
	return nil
}

// GetThirdTicketSetOpenParamList ThirdTicketSetOpenParamList Getter
func (r AlibabaDamaiMevOpenBatchpushticketAPIRequest) GetThirdTicketSetOpenParamList() []ThirdTicketPushOpenParam {
	return r._thirdTicketSetOpenParamList
}
