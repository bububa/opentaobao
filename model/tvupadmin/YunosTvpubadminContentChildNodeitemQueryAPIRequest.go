package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchildnodeitemqueryAPIRequest 查询少儿大厅类目内容 API请求
// yunos.tvpubadmin.content.child.nodeitem.query
//
// 查询少儿大厅类目内容信息
type YunostvpubadmincontentchildnodeitemqueryAPIRequest struct {
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

// NewYunostvpubadmincontentchildnodeitemqueryRequest 初始化YunostvpubadmincontentchildnodeitemqueryAPIRequest对象
func NewYunostvpubadmincontentchildnodeitemqueryRequest() *YunostvpubadmincontentchildnodeitemqueryAPIRequest {
	return &YunostvpubadmincontentchildnodeitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.nodeitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 内容名称
func (r *YunostvpubadmincontentchildnodeitemqueryAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetName() string {
	return r._name
}

// SetId is Id Setter
// 主键ID
func (r *YunostvpubadmincontentchildnodeitemqueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetId() int64 {
	return r._id
}

// SetNodeId is NodeId Setter
// 类目ID
func (r *YunostvpubadmincontentchildnodeitemqueryAPIRequest) SetNodeId(_nodeId int64) error {
	r._nodeId = _nodeId
	r.Set("node_id", _nodeId)
	return nil
}

// GetNodeId NodeId Getter
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetNodeId() int64 {
	return r._nodeId
}

// SetStatus is Status Setter
// 状态
func (r *YunostvpubadmincontentchildnodeitemqueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetStatus() int64 {
	return r._status
}

// SetPageNo is PageNo Setter
// 页码
func (r *YunostvpubadmincontentchildnodeitemqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 单页数量
func (r *YunostvpubadmincontentchildnodeitemqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmincontentchildnodeitemqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
