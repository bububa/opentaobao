package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDistributorsGetAPIRequest
获取分销商信息 API请求
taobao.fenxiao.distributors.get

查询和当前登录供应商有合作关系的分销商的信息 */
type TaobaoFenxiaoDistributorsGetAPIRequest struct {
	model.Params
	// 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
	_nicks string
}

// New
