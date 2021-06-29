package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询少儿大厅推荐内容列表 API请求
yunos.tvpubadmin.content.child.recoitem.query

查询少儿大厅推荐内容列表
*/
type YunosTvpubadminContentChildRecoitemQueryRequest struct {
    model.Params
    // 主键ID
    id   int64
    // 所属类目ID
    nodeId   int64
    // 状态
    status   int64
    // 页码
    pageNo   int64
    // 名称
    name   string
    // 单页数量
    pageSize   int64
}

// 初始化YunosTvpubadminContentChildRecoitemQueryRequest对象
func NewYunosTvpubadminContentChildRecoitemQueryRequest() *YunosTvpubadminContentChildRecoitemQueryRequest{
    return &YunosTvpubadminContentChildRecoitemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.recoitem.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 主键ID
func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetId() int64 {
    return r.id
}
// NodeId Setter
// 所属类目ID
func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetNodeId(nodeId int64) error {
    r.nodeId = nodeId
    r.Set("node_id", nodeId)
    return nil
}

// NodeId Getter
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetNodeId() int64 {
    return r.nodeId
}
// Status Setter
// 状态
func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetStatus() int64 {
    return r.status
}
// PageNo Setter
// 页码
func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// Name Setter
// 名称
func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetName() string {
    return r.name
}
// PageSize Setter
// 单页数量
func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
