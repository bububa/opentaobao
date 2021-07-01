package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaImapPvAutofillAPIRequest
属性回填接口 API请求
alibaba.imap.pv.autofill

根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息 */
type AlibabaImapPvAutofillAPIRequest struct {
	model.Params
	// 系统入参
	_topImapItemDo *TopImapItemDo
}

// New
