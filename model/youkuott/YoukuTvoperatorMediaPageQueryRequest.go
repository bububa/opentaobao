package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商全量媒资分页查询 API请求
youku.tvoperator.media.page.query

分页获取渠道全量媒资
*/
type YoukuTvoperatorMediaPageQueryRequest struct {
    model.Params
    // 系统信息（和服务提供方确认)
    _systemInfo   string
    // 从第一页开始
    _pageNo   int64
    // 页面大小
    _pageSize   int64
    // 节目programId
    _programId   int64
}

// 初始化YoukuTvoperatorMediaPageQueryRequest对象
func NewYoukuTvoperatorMediaPageQueryRequest() *YoukuTvoperatorMediaPageQueryRequest{
    return &YoukuTvoperatorMediaPageQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuTvoperatorMediaPageQueryRequest) GetApiMethodName() string {
    return "youku.tvoperator.media.page.query"
}

// IRequest interface 方法, 获取API参数
func (r YoukuTvoperatorMediaPageQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SystemInfo Setter
// 系统信息（和服务提供方确认)
func (r *YoukuTvoperatorMediaPageQueryRequest) SetSystemInfo(_systemInfo string) error {
    r._systemInfo = _systemInfo
    r.Set("system_info", _systemInfo)
    return nil
}

// SystemInfo Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetSystemInfo() string {
    return r._systemInfo
}
// PageNo Setter
// 从第一页开始
func (r *YoukuTvoperatorMediaPageQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 页面大小
func (r *YoukuTvoperatorMediaPageQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// ProgramId Setter
// 节目programId
func (r *YoukuTvoperatorMediaPageQueryRequest) SetProgramId(_programId int64) error {
    r._programId = _programId
    r.Set("program_id", _programId)
    return nil
}

// ProgramId Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetProgramId() int64 {
    return r._programId
}
