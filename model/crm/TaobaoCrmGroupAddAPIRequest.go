package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGroupAddAPIRequest
卖家创建一个分组 API请求
taobao.crm.group.add

卖家创建一个新的分组，接口返回一个创建成功的分组的id */
type TaobaoCrmGroupAddAPIRequest struct {
	model.Params
	// 分组名称，每个卖家最多可以拥有100个分组
	_groupName string
}

// New
