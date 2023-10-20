package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCreativeRptdailylistAPIRequest 创意分日数据查询 API请求
// taobao.feedflow.item.creative.rptdailylist
//
// 创意分日数据查询
type TaobaoFeedflowItemCreativeRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaoFeedflowItemCreativeRptdailylistRequest 初始化TaobaoFeedflowItemCreativeRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemCreativeRptdailylistRequest() *TaobaoFeedflowItemCreativeRptdailylistAPIRequest {
	return &TaobaoFeedflowItemCreativeRptdailylistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCreativeRptdailylistAPIRequest) Reset() {
	r._rptQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeRptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.creative.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeRptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCreativeRptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemCreativeRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaoFeedflowItemCreativeRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}

var poolTaobaoFeedflowItemCreativeRptdailylistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCreativeRptdailylistRequest()
	},
}

// GetTaobaoFeedflowItemCreativeRptdailylistRequest 从 sync.Pool 获取 TaobaoFeedflowItemCreativeRptdailylistAPIRequest
func GetTaobaoFeedflowItemCreativeRptdailylistAPIRequest() *TaobaoFeedflowItemCreativeRptdailylistAPIRequest {
	return poolTaobaoFeedflowItemCreativeRptdailylistAPIRequest.Get().(*TaobaoFeedflowItemCreativeRptdailylistAPIRequest)
}

// ReleaseTaobaoFeedflowItemCreativeRptdailylistAPIRequest 将 TaobaoFeedflowItemCreativeRptdailylistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCreativeRptdailylistAPIRequest(v *TaobaoFeedflowItemCreativeRptdailylistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCreativeRptdailylistAPIRequest.Put(v)
}
