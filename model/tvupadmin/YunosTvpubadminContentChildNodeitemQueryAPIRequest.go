package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildNodeitemQueryAPIRequest 查询少儿大厅类目内容 API请求
// yunos.tvpubadmin.content.child.nodeitem.query
//
// 查询少儿大厅类目内容信息
type YunosTvpubadminContentChildNodeitemQueryAPIRequest struct {
	model.Params
	// 内容名称
	_name string
	// 主键ID
	_id int64
	// 类目ID
	_nodeId int64
	// 状态
	_status int64
	// 页码
	_pageNo int64
	// 单页数量
	_pageSize int64
}

// NewYunosTvpubadminContentChildNodeitemQueryRequest 初始化YunosTvpubadminContentChildNodeitemQueryAPIRequest对象
func NewYunosTvpubadminContentChildNodeitemQueryRequest() *YunosTvpubadminContentChildNodeitemQueryAPIRequest {
	return &YunosTvpubadminContentChildNodeitemQueryAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) Reset() {
	r._name = ""
	r._id = 0
	r._nodeId = 0
	r._status = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.nodeitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 内容名称
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetName() string {
	return r._name
}

// SetId is Id Setter
// 主键ID
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetId() int64 {
	return r._id
}

// SetNodeId is NodeId Setter
// 类目ID
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetNodeId(_nodeId int64) error {
	r._nodeId = _nodeId
	r.Set("node_id", _nodeId)
	return nil
}

// GetNodeId NodeId Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetNodeId() int64 {
	return r._nodeId
}

// SetStatus is Status Setter
// 状态
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetStatus() int64 {
	return r._status
}

// SetPageNo is PageNo Setter
// 页码
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 单页数量
func (r *YunosTvpubadminContentChildNodeitemQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminContentChildNodeitemQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolYunosTvpubadminContentChildNodeitemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentChildNodeitemQueryRequest()
	},
}

// GetYunosTvpubadminContentChildNodeitemQueryRequest 从 sync.Pool 获取 YunosTvpubadminContentChildNodeitemQueryAPIRequest
func GetYunosTvpubadminContentChildNodeitemQueryAPIRequest() *YunosTvpubadminContentChildNodeitemQueryAPIRequest {
	return poolYunosTvpubadminContentChildNodeitemQueryAPIRequest.Get().(*YunosTvpubadminContentChildNodeitemQueryAPIRequest)
}

// ReleaseYunosTvpubadminContentChildNodeitemQueryAPIRequest 将 YunosTvpubadminContentChildNodeitemQueryAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentChildNodeitemQueryAPIRequest(v *YunosTvpubadminContentChildNodeitemQueryAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentChildNodeitemQueryAPIRequest.Put(v)
}
