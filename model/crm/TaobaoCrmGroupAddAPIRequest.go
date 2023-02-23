package crm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
