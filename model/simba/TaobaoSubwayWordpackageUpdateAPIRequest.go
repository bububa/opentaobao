package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayWordpackageUpdateAPIRequest
批量更新词包 API请求
taobao.subway.wordpackage.update

批量更新词包 */
type TaobaoSubwayWordpackageUpdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
	// 词包列表
	_wordPackageDTOS []ItemWordPackageDto
}

// New
