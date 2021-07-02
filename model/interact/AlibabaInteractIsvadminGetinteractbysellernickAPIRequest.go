package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractIsvadminGetinteractbysellernickAPIRequest 根据sellerNick获取互动实例列表 API请求
// alibaba.interact.isvadmin.getinteractbysellernick
//
// 根据sellerNick获取互动实例列表
type AlibabaInteractIsvadminGetinteractbysellernickAPIRequest struct {
	model.Params
}

// NewAlibabaInteractIsvadminGetinteractbysellernickRequest 初始化AlibabaInteractIsvadminGetinteractbysellernickAPIRequest对象
func NewAlibabaInteractIsvadminGetinteractbysellernickRequest() *AlibabaInteractIsvadminGetinteractbysellernickAPIRequest {
	return &AlibabaInteractIsvadminGetinteractbysellernickAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminGetinteractbysellernickAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.isvadmin.getinteractbysellernick"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminGetinteractbysellernickAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
