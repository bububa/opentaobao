package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchildleafnodegetAPIRequest 获取少儿大厅二级类目列表 API请求
// yunos.tvpubadmin.content.child.leafnode.get
//
// 获取少儿大厅二级类目列表的接口
type YunostvpubadmincontentchildleafnodegetAPIRequest struct {
	model.Params
}

// NewYunostvpubadmincontentchildleafnodegetRequest 初始化YunostvpubadmincontentchildleafnodegetAPIRequest对象
func NewYunostvpubadmincontentchildleafnodegetRequest() *YunostvpubadmincontentchildleafnodegetAPIRequest {
	return &YunostvpubadmincontentchildleafnodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentchildleafnodegetAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.leafnode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentchildleafnodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentchildleafnodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
