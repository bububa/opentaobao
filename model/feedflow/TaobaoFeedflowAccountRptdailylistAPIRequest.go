package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowAccountRptdailylistAPIRequest 获取广告主分日数据 API请求
// taobao.feedflow.account.rptdailylist
//
// 获取广告主分日数据
type TaobaoFeedflowAccountRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaoFeedflowAccountRptdailylistRequest 初始化TaobaoFeedflowAccountRptdailylistAPIRequest对象
func NewTaobaoFeedflowAccountRptdailylistRequest() *TaobaoFeedflowAccountRptdailylistAPIRequest {
	return &TaobaoFeedflowAccountRptdailylistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowAccountRptdailylistAPIRequest) Reset() {
	r._rptQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowAccountRptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.account.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowAccountRptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowAccountRptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowAccountRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaoFeedflowAccountRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}

var poolTaobaoFeedflowAccountRptdailylistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowAccountRptdailylistRequest()
	},
}

// GetTaobaoFeedflowAccountRptdailylistRequest 从 sync.Pool 获取 TaobaoFeedflowAccountRptdailylistAPIRequest
func GetTaobaoFeedflowAccountRptdailylistAPIRequest() *TaobaoFeedflowAccountRptdailylistAPIRequest {
	return poolTaobaoFeedflowAccountRptdailylistAPIRequest.Get().(*TaobaoFeedflowAccountRptdailylistAPIRequest)
}

// ReleaseTaobaoFeedflowAccountRptdailylistAPIRequest 将 TaobaoFeedflowAccountRptdailylistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowAccountRptdailylistAPIRequest(v *TaobaoFeedflowAccountRptdailylistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowAccountRptdailylistAPIRequest.Put(v)
}
