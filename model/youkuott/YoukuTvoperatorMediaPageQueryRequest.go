package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商全量媒资分页查询 APIRequest
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

func NewYoukuTvoperatorMediaPageQueryRequest() *YoukuTvoperatorMediaPageQueryRequest{
    return &YoukuTvoperatorMediaPageQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuTvoperatorMediaPageQueryRequest) GetApiMethodName() string {
    return "youku.tvoperator.media.page.query"
}

func (r YoukuTvoperatorMediaPageQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuTvoperatorMediaPageQueryRequest) SetSystemInfo(systemInfo string) error {
    r.systemInfo = systemInfo
    r.Set("system_info", systemInfo)
    return nil
}

func (r YoukuTvoperatorMediaPageQueryRequest) GetSystemInfo() string {
    return r.systemInfo
}

func (r *YoukuTvoperatorMediaPageQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YoukuTvoperatorMediaPageQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YoukuTvoperatorMediaPageQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YoukuTvoperatorMediaPageQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *YoukuTvoperatorMediaPageQueryRequest) SetProgramId(programId int64) error {
    r.programId = programId
    r.Set("program_id", programId)
    return nil
}

func (r YoukuTvoperatorMediaPageQueryRequest) GetProgramId() int64 {
    return r.programId
}

