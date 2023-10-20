package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupRptdailylistAPIRequest 推广单元分日数据查询 API请求
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
type TaobaoFeedflowItemAdgroupRptdailylistAPIRequest struct {
	model.Params
	// 查询条件
	_rptQueryDTO *RptQueryDto
}

// NewTaobaoFeedflowItemAdgroupRptdailylistRequest 初始化TaobaoFeedflowItemAdgroupRptdailylistAPIRequest对象
func NewTaobaoFeedflowItemAdgroupRptdailylistRequest() *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest {
	return &TaobaoFeedflowItemAdgroupRptdailylistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) Reset() {
	r._rptQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.rptdailylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRptQueryDTO is RptQueryDTO Setter
// 查询条件
func (r *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) SetRptQueryDTO(_rptQueryDTO *RptQueryDto) error {
	r._rptQueryDTO = _rptQueryDTO
	r.Set("rpt_query_d_t_o", _rptQueryDTO)
	return nil
}

// GetRptQueryDTO RptQueryDTO Getter
func (r TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) GetRptQueryDTO() *RptQueryDto {
	return r._rptQueryDTO
}

var poolTaobaoFeedflowItemAdgroupRptdailylistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupRptdailylistRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupRptdailylistRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupRptdailylistAPIRequest
func GetTaobaoFeedflowItemAdgroupRptdailylistAPIRequest() *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest {
	return poolTaobaoFeedflowItemAdgroupRptdailylistAPIRequest.Get().(*TaobaoFeedflowItemAdgroupRptdailylistAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupRptdailylistAPIRequest 将 TaobaoFeedflowItemAdgroupRptdailylistAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupRptdailylistAPIRequest(v *TaobaoFeedflowItemAdgroupRptdailylistAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupRptdailylistAPIRequest.Put(v)
}
