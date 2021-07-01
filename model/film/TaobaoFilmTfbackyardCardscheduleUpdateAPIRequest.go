package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest
CGV影城卡排期数据传输 API请求
taobao.film.tfbackyard.cardschedule.update

cgv影城卡排期价格数据传输API */
type TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest struct {
	model.Params
	// CGV影城卡价格数据
	_jsonData string
}

// New
