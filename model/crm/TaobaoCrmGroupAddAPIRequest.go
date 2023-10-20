package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupAddAPIRequest 卖家创建一个分组 API请求
// taobao.crm.group.add
//
// 卖家创建一个新的分组，接口返回一个创建成功的分组的id
type TaobaoCrmGroupAddAPIRequest struct {
	model.Params
	// 分组名称，每个卖家最多可以拥有100个分组
	_groupName string
}

// NewTaobaoCrmGroupAddRequest 初始化TaobaoCrmGroupAddAPIRequest对象
func NewTaobaoCrmGroupAddRequest() *TaobaoCrmGroupAddAPIRequest {
	return &TaobaoCrmGroupAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGroupAddAPIRequest) Reset() {
	r._groupName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupAddAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGroupAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupName is GroupName Setter
// 分组名称，每个卖家最多可以拥有100个分组
func (r *TaobaoCrmGroupAddAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaoCrmGroupAddAPIRequest) GetGroupName() string {
	return r._groupName
}

var poolTaobaoCrmGroupAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGroupAddRequest()
	},
}

// GetTaobaoCrmGroupAddRequest 从 sync.Pool 获取 TaobaoCrmGroupAddAPIRequest
func GetTaobaoCrmGroupAddAPIRequest() *TaobaoCrmGroupAddAPIRequest {
	return poolTaobaoCrmGroupAddAPIRequest.Get().(*TaobaoCrmGroupAddAPIRequest)
}

// ReleaseTaobaoCrmGroupAddAPIRequest 将 TaobaoCrmGroupAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGroupAddAPIRequest(v *TaobaoCrmGroupAddAPIRequest) {
	v.Reset()
	poolTaobaoCrmGroupAddAPIRequest.Put(v)
}
