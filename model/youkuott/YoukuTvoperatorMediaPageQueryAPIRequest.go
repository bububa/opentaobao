package youkuott

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuTvoperatorMediaPageQueryAPIRequest 运营商全量媒资分页查询 API请求
// youku.tvoperator.media.page.query
//
// 分页获取渠道全量媒资
type YoukuTvoperatorMediaPageQueryAPIRequest struct {
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

// NewYoukuTvoperatorMediaPageQueryRequest 初始化YoukuTvoperatorMediaPageQueryAPIRequest对象
func NewYoukuTvoperatorMediaPageQueryRequest() *YoukuTvoperatorMediaPageQueryAPIRequest {
	return &YoukuTvoperatorMediaPageQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuTvoperatorMediaPageQueryAPIRequest) Reset() {
	r._systemInfo = ""
	r._pageNo = 0
	r._pageSize = 0
	r._programId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuTvoperatorMediaPageQueryAPIRequest) GetApiMethodName() string {
	return "youku.tvoperator.media.page.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuTvoperatorMediaPageQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuTvoperatorMediaPageQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSystemInfo is SystemInfo Setter
// 系统信息（和服务提供方确认)
func (r *YoukuTvoperatorMediaPageQueryAPIRequest) SetSystemInfo(_systemInfo string) error {
	r._systemInfo = _systemInfo
	r.Set("system_info", _systemInfo)
	return nil
}

// GetSystemInfo SystemInfo Getter
func (r YoukuTvoperatorMediaPageQueryAPIRequest) GetSystemInfo() string {
	return r._systemInfo
}

// SetPageNo is PageNo Setter
// 从第一页开始
func (r *YoukuTvoperatorMediaPageQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YoukuTvoperatorMediaPageQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *YoukuTvoperatorMediaPageQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YoukuTvoperatorMediaPageQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetProgramId is ProgramId Setter
// 节目programId
func (r *YoukuTvoperatorMediaPageQueryAPIRequest) SetProgramId(_programId int64) error {
	r._programId = _programId
	r.Set("program_id", _programId)
	return nil
}

// GetProgramId ProgramId Getter
func (r YoukuTvoperatorMediaPageQueryAPIRequest) GetProgramId() int64 {
	return r._programId
}

var poolYoukuTvoperatorMediaPageQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuTvoperatorMediaPageQueryRequest()
	},
}

// GetYoukuTvoperatorMediaPageQueryRequest 从 sync.Pool 获取 YoukuTvoperatorMediaPageQueryAPIRequest
func GetYoukuTvoperatorMediaPageQueryAPIRequest() *YoukuTvoperatorMediaPageQueryAPIRequest {
	return poolYoukuTvoperatorMediaPageQueryAPIRequest.Get().(*YoukuTvoperatorMediaPageQueryAPIRequest)
}

// ReleaseYoukuTvoperatorMediaPageQueryAPIRequest 将 YoukuTvoperatorMediaPageQueryAPIRequest 放入 sync.Pool
func ReleaseYoukuTvoperatorMediaPageQueryAPIRequest(v *YoukuTvoperatorMediaPageQueryAPIRequest) {
	v.Reset()
	poolYoukuTvoperatorMediaPageQueryAPIRequest.Put(v)
}
