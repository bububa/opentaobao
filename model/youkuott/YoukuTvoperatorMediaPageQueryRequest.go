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
    systemInfo   string
    // 从第一页开始
    pageNo   int64
    // 页面大小
    pageSize   int64
    // 节目programId
    programId   int64
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
func (r *YoukuTvoperatorMediaPageQueryRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

// SystemInfo Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetSystemInfo() string {
    return r.systemInfo
}
// PageNo Setter
// 从第一页开始
func (r *YoukuTvoperatorMediaPageQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 页面大小
func (r *YoukuTvoperatorMediaPageQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// ProgramId Setter
// 节目programId
func (r *YoukuTvoperatorMediaPageQueryRequest) SetProgramId(programId int64) error {
    r.programId = programId
    r.Set("program_id", programId)
    return nil
}

// ProgramId Getter
func (r YoukuTvoperatorMediaPageQueryRequest) GetProgramId() int64 {
    return r.programId
}
