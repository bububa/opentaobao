package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayCiaUpdateAPIRequest
批量修改单元智能出价 API请求
taobao.subway.cia.update

批量修改直通车推广单元的智能出价配置 */
type TaobaoSubwayCiaUpdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 系统自动生成
	_ciaConfigs []CiaUpdateDto
}

// New
