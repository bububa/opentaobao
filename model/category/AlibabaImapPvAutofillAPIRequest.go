package category

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaImapPvAutofillAPIRequest 属性回填接口 API请求
// alibaba.imap.pv.autofill
//
// 根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
type AlibabaImapPvAutofillAPIRequest struct {
	model.Params
	// 系统入参
	_topImapItemDo *TopImapItemDo
}

// NewAlibabaImapPvAutofillRequest 初始化AlibabaImapPvAutofillAPIRequest对象
func NewAlibabaImapPvAutofillRequest() *AlibabaImapPvAutofillAPIRequest {
	return &AlibabaImapPvAutofillAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaImapPvAutofillAPIRequest) Reset() {
	r._topImapItemDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaImapPvAutofillAPIRequest) GetApiMethodName() string {
	return "alibaba.imap.pv.autofill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaImapPvAutofillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaImapPvAutofillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopImapItemDo is TopImapItemDo Setter
// 系统入参
func (r *AlibabaImapPvAutofillAPIRequest) SetTopImapItemDo(_topImapItemDo *TopImapItemDo) error {
	r._topImapItemDo = _topImapItemDo
	r.Set("top_imap_item_do", _topImapItemDo)
	return nil
}

// GetTopImapItemDo TopImapItemDo Getter
func (r AlibabaImapPvAutofillAPIRequest) GetTopImapItemDo() *TopImapItemDo {
	return r._topImapItemDo
}

var poolAlibabaImapPvAutofillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaImapPvAutofillRequest()
	},
}

// GetAlibabaImapPvAutofillRequest 从 sync.Pool 获取 AlibabaImapPvAutofillAPIRequest
func GetAlibabaImapPvAutofillAPIRequest() *AlibabaImapPvAutofillAPIRequest {
	return poolAlibabaImapPvAutofillAPIRequest.Get().(*AlibabaImapPvAutofillAPIRequest)
}

// ReleaseAlibabaImapPvAutofillAPIRequest 将 AlibabaImapPvAutofillAPIRequest 放入 sync.Pool
func ReleaseAlibabaImapPvAutofillAPIRequest(v *AlibabaImapPvAutofillAPIRequest) {
	v.Reset()
	poolAlibabaImapPvAutofillAPIRequest.Put(v)
}
