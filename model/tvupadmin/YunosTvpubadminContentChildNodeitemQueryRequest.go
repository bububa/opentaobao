package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询少儿大厅类目内容 API请求
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

// 初始化YunosTvpubadminContentChildNodeitemQueryRequest对象
func NewYunosTvpubadminContentChildNodeitemQueryRequest() *YunosTvpubadminContentChildNodeitemQueryRequest{
    return &YunosTvpubadminContentChildNodeitemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.nodeitem.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 主键ID
func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetId() int64 {
    return r.id
}
// NodeId Setter
// 类目ID
func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetNodeId(nodeId int64) error {
    r.nodeId = nodeId
    r.Set("node_id", nodeId)
    return nil
}

// NodeId Getter
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetNodeId() int64 {
    return r.nodeId
}
// Status Setter
// 状态
func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetStatus() int64 {
    return r.status
}
// PageNo Setter
// 页码
func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// Name Setter
// 内容名称
func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetName() string {
    return r.name
}
// PageSize Setter
// 单页数量
func (r *YunosTvpubadminContentChildNodeitemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentChildNodeitemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
