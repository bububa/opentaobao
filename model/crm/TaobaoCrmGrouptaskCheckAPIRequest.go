package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrouptaskCheckAPIRequest 查询分组任务是否完成 API请求
// taobao.crm.grouptask.check
//
// 检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组&lt;br/&gt;若分组上有任务则该属性不能被操作。
type TaobaoCrmGrouptaskCheckAPIRequest struct {
	model.Params
	// 分组id
	_groupId int64
}

// NewTaobaoCrmGrouptaskCheckRequest 初始化TaobaoCrmGrouptaskCheckAPIRequest对象
func NewTaobaoCrmGrouptaskCheckRequest() *TaobaoCrmGrouptaskCheckAPIRequest {
	return &TaobaoCrmGrouptaskCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGrouptaskCheckAPIRequest) Reset() {
	r._groupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrouptaskCheckAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grouptask.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrouptaskCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGrouptaskCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 分组id
func (r *TaobaoCrmGrouptaskCheckAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r TaobaoCrmGrouptaskCheckAPIRequest) GetGroupId() int64 {
	return r._groupId
}

var poolTaobaoCrmGrouptaskCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGrouptaskCheckRequest()
	},
}

// GetTaobaoCrmGrouptaskCheckRequest 从 sync.Pool 获取 TaobaoCrmGrouptaskCheckAPIRequest
func GetTaobaoCrmGrouptaskCheckAPIRequest() *TaobaoCrmGrouptaskCheckAPIRequest {
	return poolTaobaoCrmGrouptaskCheckAPIRequest.Get().(*TaobaoCrmGrouptaskCheckAPIRequest)
}

// ReleaseTaobaoCrmGrouptaskCheckAPIRequest 将 TaobaoCrmGrouptaskCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGrouptaskCheckAPIRequest(v *TaobaoCrmGrouptaskCheckAPIRequest) {
	v.Reset()
	poolTaobaoCrmGrouptaskCheckAPIRequest.Put(v)
}
