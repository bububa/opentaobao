package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchildrecoitemofflineAPIRequest 下线少儿推荐内容接口 API请求
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
type YunostvpubadmincontentchildrecoitemofflineAPIRequest struct {
	model.Params
	// 推荐内容ID
	_recItemId int64
}

// NewYunostvpubadmincontentchildrecoitemofflineRequest 初始化YunostvpubadmincontentchildrecoitemofflineAPIRequest对象
func NewYunostvpubadmincontentchildrecoitemofflineRequest() *YunostvpubadmincontentchildrecoitemofflineAPIRequest {
	return &YunostvpubadmincontentchildrecoitemofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentchildrecoitemofflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.recoitem.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentchildrecoitemofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentchildrecoitemofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecItemId is RecItemId Setter
// 推荐内容ID
func (r *YunostvpubadmincontentchildrecoitemofflineAPIRequest) SetRecItemId(_recItemId int64) error {
	r._recItemId = _recItemId
	r.Set("rec_item_id", _recItemId)
	return nil
}

// GetRecItemId RecItemId Getter
func (r YunostvpubadmincontentchildrecoitemofflineAPIRequest) GetRecItemId() int64 {
	return r._recItemId
}
