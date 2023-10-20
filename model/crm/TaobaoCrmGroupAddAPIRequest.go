package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmgroupaddAPIRequest 卖家创建一个分组 API请求
// taobao.crm.group.add
//
// 卖家创建一个新的分组，接口返回一个创建成功的分组的id
type TaobaocrmgroupaddAPIRequest struct {
	model.Params
	// 分组名称，每个卖家最多可以拥有100个分组
	_groupName string
}

// NewTaobaocrmgroupaddRequest 初始化TaobaocrmgroupaddAPIRequest对象
func NewTaobaocrmgroupaddRequest() *TaobaocrmgroupaddAPIRequest {
	return &TaobaocrmgroupaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmgroupaddAPIRequest) GetApiMethodName() string {
	return "taobao.crm.group.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmgroupaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmgroupaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupName is GroupName Setter
// 分组名称，每个卖家最多可以拥有100个分组
func (r *TaobaocrmgroupaddAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaocrmgroupaddAPIRequest) GetGroupName() string {
	return r._groupName
}
