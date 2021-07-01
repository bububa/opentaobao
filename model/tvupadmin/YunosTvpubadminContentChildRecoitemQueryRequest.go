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
type YunosTvpubadminContentChildRecoitemQueryAPIRequest struct {
    model.Params
    // 主键ID
    _id   int64
    // 所属类目ID
    _nodeId   int64
    // 状态
    _status   int64
    // 页码
    _pageNo   int64
    // 名称
    _name   string
    // 单页数量
    _pageSize   int64
}

// 初始化YunosTvpubadminContentChildRecoitemQueryAPIRequest对象
func NewYunosTvpubadminContentChildRecoitemQueryRequest() *YunosTvpubadminContentChildRecoitemQueryAPIRequest{
    return &YunosTvpubadminContentChildRecoitemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.recoitem.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 主键ID
func (r *YunosTvpubadminContentChildRecoitemQueryAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetId() int64 {
    return r._id
}
// NodeId Setter
// 所属类目ID
func (r *YunosTvpubadminContentChildRecoitemQueryAPIRequest) SetNodeId(_nodeId int64) error {
    r._nodeId = _nodeId
    r.Set("node_id", _nodeId)
    return nil
}

// NodeId Getter
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetNodeId() int64 {
    return r._nodeId
}
// Status Setter
// 状态
func (r *YunosTvpubadminContentChildRecoitemQueryAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetStatus() int64 {
    return r._status
}
// PageNo Setter
// 页码
func (r *YunosTvpubadminContentChildRecoitemQueryAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// Name Setter
// 名称
func (r *YunosTvpubadminContentChildRecoitemQueryAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetName() string {
    return r._name
}
// PageSize Setter
// 单页数量
func (r *YunosTvpubadminContentChildRecoitemQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentChildRecoitemQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
