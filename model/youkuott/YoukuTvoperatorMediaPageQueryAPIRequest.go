package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukutvoperatormediapagequeryAPIRequest 运营商全量媒资分页查询 API请求
// youku.tvoperator.media.page.query
//
// 分页获取渠道全量媒资
type YoukutvoperatormediapagequeryAPIRequest struct {
	model.Params
	// 系统信息（和服务提供方确认)
	_systemInfo string
	// 从第一页开始
	_pageNo int64
	// 页面大小
	_pageSize int64
	// 节目programId
	_programId int64
}

// NewYoukutvoperatormediapagequeryRequest 初始化YoukutvoperatormediapagequeryAPIRequest对象
func NewYoukutvoperatormediapagequeryRequest() *YoukutvoperatormediapagequeryAPIRequest {
	return &YoukutvoperatormediapagequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukutvoperatormediapagequeryAPIRequest) GetApiMethodName() string {
	return "youku.tvoperator.media.page.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukutvoperatormediapagequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukutvoperatormediapagequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息（和服务提供方确认)
func (r *YoukutvoperatormediapagequeryAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r YoukutvoperatormediapagequeryAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPageNo is PageNo Setter
// 从第一页开始
func (r *YoukutvoperatormediapagequeryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YoukutvoperatormediapagequeryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *YoukutvoperatormediapagequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YoukutvoperatormediapagequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetProgramId is ProgramId Setter
// 节目programId
func (r *YoukutvoperatormediapagequeryAPIRequest) SetProgramId(_programId int64) error {
	r._programId = _programId
	r.Set("program_id", _programId)
	return nil
}

// GetProgramId ProgramId Getter
func (r YoukutvoperatormediapagequeryAPIRequest) GetProgramId() int64 {
	return r._programId
}
