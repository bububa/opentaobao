package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// TaobaoFilmTfbackyardCardscheduleUpdate CGV影城卡排期数据传输
// taobao.film.tfbackyard.cardschedule.update
//
// cgv影城卡排期价格数据传输API
func TaobaoFilmTfbackyardCardscheduleUpdate(clt *core.SDKClient, req *film.TaobaoFilmTfbackyardCardscheduleUpdateAPIRequest, resp *film.TaobaoFilmTfbackyardCardscheduleUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
