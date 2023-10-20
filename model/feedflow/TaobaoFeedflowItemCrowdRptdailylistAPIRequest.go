package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdRptdailylistAPIRequest 定向分日数据查询 API请求
// taobao.feedflow.item.crowd.rptdailylist
//
// 定向分日数据查询
type TaobaoFeedflowItemCrowdRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaoFeedflowItemCrowdRptdailylistRequest 初始化TaobaoFeedflowItemCrowdRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemCrowdRptdailylistRequest() *TaobaoFeedflowItemCrowdRptdailylistAPIRequest {
	return &TaobaoFeedflowItemCrowdRptdailylistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCrowdRptdailylistAPIRequest) Reset() {
	r._rptQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdRptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdRptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCrowdRptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCrowdRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaoFeedflowItemCrowdRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}

var poolTaobaoFeedflowItemCrowdRptdailylistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCrowdRptdailylistRequest()
	},
}

// GetTaobaoFeedflowItemCrowdRptdailylistRequest 从 sync.Pool 获取 TaobaoFeedflowItemCrowdRptdailylistAPIRequest
func GetTaobaoFeedflowItemCrowdRptdailylistAPIRequest() *TaobaoFeedflowItemCrowdRptdailylistAPIRequest {
	return poolTaobaoFeedflowItemCrowdRptdailylistAPIRequest.Get().(*TaobaoFeedflowItemCrowdRptdailylistAPIRequest)
}

// ReleaseTaobaoFeedflowItemCrowdRptdailylistAPIRequest 将 TaobaoFeedflowItemCrowdRptdailylistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdRptdailylistAPIRequest(v *TaobaoFeedflowItemCrowdRptdailylistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdRptdailylistAPIRequest.Put(v)
}
