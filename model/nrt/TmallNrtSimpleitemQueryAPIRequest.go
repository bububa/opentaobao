package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtSimpleitemQueryAPIRequest
简易商品查询接口 API请求
tmall.nrt.simpleitem.query

为居然之家和阿里的合资公司 homeStyler提供简易的商品信息查询 包含商品名称  图片 状态

后续合资公司服务会迁移到内网 暂时过渡用 */
type TmallNrtSimpleitemQueryAPIRequest struct {
	model.Params
	// 商品编码数组
	_ids []int64
}

// New
