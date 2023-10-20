package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocrmmemberinfoupdateAPIRequest 编辑会员资料 API请求
// taobao.crm.memberinfo.update
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
type TaobaocrmmemberinfoupdateAPIRequest struct {
	model.Params
}

// NewTaobaocrmmemberinfoupdateRequest 初始化TaobaocrmmemberinfoupdateAPIRequest对象
func NewTaobaocrmmemberinfoupdateRequest() *TaobaocrmmemberinfoupdateAPIRequest {
	return &TaobaocrmmemberinfoupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocrmmemberinfoupdateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.memberinfo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocrmmemberinfoupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocrmmemberinfoupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
