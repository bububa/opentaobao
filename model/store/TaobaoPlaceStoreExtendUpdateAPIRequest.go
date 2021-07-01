package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreExtendUpdateAPIRequest
商户门店拓展信息更新接口 API请求
taobao.place.store.extend.update

更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口 */
type TaobaoPlaceStoreExtendUpdateAPIRequest struct {
	model.Params
	// 更新数据
	_paramUpdateStoreExtendDTO *UpdateStoreExtendDto
}

// New
