package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscToolCheckAPIRequest UMP工具检测 API请求
// taobao.promotionmisc.tool.check
//
// UMP工具检测。ISV通过该接口检测（通过taobao.ump.tool.add）创建的UMP工具（tool）是否符合规范，如果不符合，则返回错误信息和对应的解决方案的；工具检测通过后才可以提交工具审核邮件，提交工具审核时，需提供该接口的返回值。
type TaobaoPromotionmiscToolCheckAPIRequest struct {
	model.Params
	// 可使用的元数据。PRD审核后，会告诉isv可使用的元数据。
	_metaAllow string
	// 工具ID, taobao.ump.tool.add成功后返回的id。
	_toolId int64
}

// NewTaobaoPromotionmiscToolCheckRequest 初始化TaobaoPromotionmiscToolCheckAPIRequest对象
func NewTaobaoPromotionmiscToolCheckRequest() *TaobaoPromotionmiscToolCheckAPIRequest {
	return &TaobaoPromotionmiscToolCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscToolCheckAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.tool.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscToolCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMetaAllow is MetaAllow Setter
// 可使用的元数据。PRD审核后，会告诉isv可使用的元数据。
func (r *TaobaoPromotionmiscToolCheckAPIRequest) SetMetaAllow(_metaAllow string) error {
	r._metaAllow = _metaAllow
	r.Set("meta_allow", _metaAllow)
	return nil
}

// GetMetaAllow MetaAllow Getter
func (r TaobaoPromotionmiscToolCheckAPIRequest) GetMetaAllow() string {
	return r._metaAllow
}

// SetToolId is ToolId Setter
// 工具ID, taobao.ump.tool.add成功后返回的id。
func (r *TaobaoPromotionmiscToolCheckAPIRequest) SetToolId(_toolId int64) error {
	r._toolId = _toolId
	r.Set("tool_id", _toolId)
	return nil
}

// GetToolId ToolId Getter
func (r TaobaoPromotionmiscToolCheckAPIRequest) GetToolId() int64 {
	return r._toolId
}
