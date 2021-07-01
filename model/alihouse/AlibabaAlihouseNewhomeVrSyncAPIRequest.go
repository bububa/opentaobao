package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeVrSyncAPIRequest
VR关系数据同步 API请求
alibaba.alihouse.newhome.vr.sync

对接易居VR关系数据迁移 */
type AlibabaAlihouseNewhomeVrSyncAPIRequest struct {
	model.Params
	// VR提取码
	_extractedCode string
	// 户型ID
	_layoutInfoId string
	// 是否生效
	_isValid string
	// 封面图
	_coverImage string
	// vr展示链接
	_vrUrl string
	// 营销户型图
	_markingLayoutImg string
	// 是否是单层 1 是 0 否
	_isSingleLayout int64
}

// New
