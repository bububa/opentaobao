package category

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// AlibabaImapPvAutofill 属性回填接口
// alibaba.imap.pv.autofill
//
// 根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
func AlibabaImapPvAutofill(ctx context.Context, clt *core.SDKClient, req *category.AlibabaImapPvAutofillAPIRequest, resp *category.AlibabaImapPvAutofillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
