package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiAlbumDetailGetAPIRequest
虾米音乐专辑详情接口 API请求
alibaba.xiami.api.album.detail.get

虾米音乐专辑详情接口 */
type AlibabaXiamiApiAlbumDetailGetAPIRequest struct {
	model.Params
	// 专辑ID
	_id int64
	// 是否获取完整描述
	_fullDes bool
}

// New
