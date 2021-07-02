package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildRecoitemOfflineAPIRequest 下线少儿推荐内容接口 API请求
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
type YunosTvpubadminContentChildRecoitemOfflineAPIRequest struct {
	model.Params
	// 推荐内容ID
	_recItemId int64
}

// NewYunosTvpubadminContentChildRecoitemOfflineRequest 初始化YunosTvpubadminContentChildRecoitemOfflineAPIRequest对象
func NewYunosTvpubadminContentChildRecoitemOfflineRequest() *YunosTvpubadminContentChildRecoitemOfflineAPIRequest {
	return &YunosTvpubadminContentChildRecoitemOfflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildRecoitemOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.recoitem.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildRecoitemOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RecItemId Setter
// 推荐内容ID
func (r *YunosTvpubadminContentChildRecoitemOfflineAPIRequest) SetRecItemId(_recItemId int64) error {
	r._recItemId = _recItemId
	r.Set("rec_item_id", _recItemId)
	return nil
}

// Get RecItemId Getter
func (r YunosTvpubadminContentChildRecoitemOfflineAPIRequest) GetRecItemId() int64 {
	return r._recItemId
}
