package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductcatsGetAPIRequest
查询产品线列表 API请求
taobao.fenxiao.productcats.get

查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参 */
type TaobaoFenxiaoProductcatsGetAPIRequest struct {
	model.Params
	// 返回字段列表
	_fields string
}

// New
