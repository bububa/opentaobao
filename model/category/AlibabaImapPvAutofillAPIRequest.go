package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaimappvautofillAPIRequest 属性回填接口 API请求
// alibaba.imap.pv.autofill
//
// 根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
type AlibabaimappvautofillAPIRequest struct {
	model.Params
	// 系统入参
	_topImapItemDo *TopImapItemDo
}

// NewAlibabaimappvautofillRequest 初始化AlibabaimappvautofillAPIRequest对象
func NewAlibabaimappvautofillRequest() *AlibabaimappvautofillAPIRequest {
	return &AlibabaimappvautofillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaimappvautofillAPIRequest) GetApiMethodName() string {
	return "alibaba.imap.pv.autofill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaimappvautofillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaimappvautofillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopImapItemDo is TopImapItemDo Setter
// 系统入参
func (r *AlibabaimappvautofillAPIRequest) SetTopImapItemDo(_topImapItemDo *TopImapItemDo) error {
	r._topImapItemDo = _topImapItemDo
	r.Set("top_imap_item_do", _topImapItemDo)
	return nil
}

// GetTopImapItemDo TopImapItemDo Getter
func (r AlibabaimappvautofillAPIRequest) GetTopImapItemDo() *TopImapItemDo {
	return r._topImapItemDo
}
