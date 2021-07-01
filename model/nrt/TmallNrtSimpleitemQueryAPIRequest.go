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

// NewTmallNrtSimpleitemQueryRequest 初始化TmallNrtSimpleitemQueryAPIRequest对象
func NewTmallNrtSimpleitemQueryRequest() *TmallNrtSimpleitemQueryAPIRequest {
	return &TmallNrtSimpleitemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtSimpleitemQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.simpleitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtSimpleitemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Ids Setter
// 商品编码数组
func (r *TmallNrtSimpleitemQueryAPIRequest) SetIds(_ids []int64) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// Get Ids Getter
func (r TmallNrtSimpleitemQueryAPIRequest) GetIds() []int64 {
	return r._ids
}
