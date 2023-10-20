package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest 创建异步下载任务 API请求
// taobao.universalbp.report.async.create.download.task
//
// 入参报表查询信息，出参下载任务id
type TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// reportQueryVO
	_reportQueryVO *ReportQueryVo
}

// NewTaobaoUniversalbpReportAsyncCreateDownloadTaskRequest 初始化TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest对象
func NewTaobaoUniversalbpReportAsyncCreateDownloadTaskRequest() *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest {
	return &TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) Reset() {
	r._topServiceContext = nil
	r._reportQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.report.async.create.download.task"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetReportQueryVO is ReportQueryVO Setter
// reportQueryVO
func (r *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) SetReportQueryVO(_reportQueryVO *ReportQueryVo) error {
	r._reportQueryVO = _reportQueryVO
	r.Set("report_query_v_o", _reportQueryVO)
	return nil
}

// GetReportQueryVO ReportQueryVO Getter
func (r TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) GetReportQueryVO() *ReportQueryVo {
	return r._reportQueryVO
}

var poolTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpReportAsyncCreateDownloadTaskRequest()
	},
}

// GetTaobaoUniversalbpReportAsyncCreateDownloadTaskRequest 从 sync.Pool 获取 TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest
func GetTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest() *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest {
	return poolTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest.Get().(*TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest)
}

// ReleaseTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest 将 TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest(v *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIRequest.Put(v)
}
