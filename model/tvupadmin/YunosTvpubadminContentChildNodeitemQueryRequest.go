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
type YunosTvpubadminContentChildNodeitemQueryAPIRequest struct {
    model.Params
    // 主键ID
    _id   int64
    // 类目ID
    _nodeId   int64
    // 状态
    _status   int64
    // 页码
    _pageNo   int64
    // 内容名称
    _name   string
    // 单页数量
    _pageSize   int64
}

// 初始化YunosTvpubadminContentChildNodeitemQueryAPIRequest对象
func NewYunosTvpubadminContentChildNodeitemQueryRequest() *YunosTvpubadminContentChildNodeitemQueryAPIRequest{
    return &YunosTvpubadminContentChildNodeitemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.child.nodeitem.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 主键ID
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetId() int64 {
    return r._id
}
// NodeId Setter
// 类目ID
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetNodeId(_nodeId int64) error {
    r._nodeId = _nodeId
    r.Set("node_id", _nodeId)
    return nil
}

// NodeId Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetNodeId() int64 {
    return r._nodeId
}
// Status Setter
// 状态
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetStatus() int64 {
    return r._status
}
// PageNo Setter
// 页码
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// Name Setter
// 内容名称
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetName() string {
    return r._name
}
// PageSize Setter
// 单页数量
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
