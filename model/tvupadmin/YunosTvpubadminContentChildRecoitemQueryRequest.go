package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询少儿大厅推荐内容列表 APIRequest
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

func NewYunosTvpubadminContentChildRecoitemQueryRequest() *YunosTvpubadminContentChildRecoitemQueryRequest{
    return &YunosTvpubadminContentChildRecoitemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.recoitem.query"
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetId() int64 {
    return r.id
}

func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetNodeId(nodeId int64) error {
    r.nodeId = nodeId
    r.Set("node_id", nodeId)
    return nil
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetNodeId() int64 {
    return r.nodeId
}

func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetStatus() int64 {
    return r.status
}

func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetName() string {
    return r.name
}

func (r *YunosTvpubadminContentChildRecoitemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminContentChildRecoitemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

