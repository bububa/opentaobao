package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询少儿大厅类目内容 APIRequest
yunos.tvpubadmin.content.child.nodeitem.query

查询少儿大厅类目内容信息
*/
type YunosTvpubadminContentChildNodeitemQueryRequest struct {
    model.Params

    // 主键ID
    id   int64 

    // 类目ID
    nodeId   int64 

    // 状态
    status   int64 

    // 页码
    pageNo   int64 

    // 内容名称
    name   string 

    // 单页数量
    pageSize   int64 

}

func NewYunosTvpubadminContentChildNodeitemQueryRequest() *YunosTvpubadminContentChildNodeitemQueryRequest{
    return &YunosTvpubadminContentChildNodeitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.nodeitem.query"
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetId() int64 {
    return r.id
}

func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetNodeId(nodeId int64) error {
    r.nodeId = nodeId
    r.Set("node_id", nodeId)
    return nil
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetNodeId() int64 {
    return r.nodeId
}

func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetStatus() int64 {
    return r.status
}

func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetName() string {
    return r.name
}

func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

