package film

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/film"
)

// Taobaofilmtfbackyardcardscheduleupdate CGV影城卡排期数据传输
// taobao.film.tfbackyard.cardschedule.update
//
// cgv影城卡排期价格数据传输API
func Taobaofilmtfbackyardcardscheduleupdate(clt *core.SDKClient, req *film.TaobaofilmtfbackyardcardscheduleupdateAPIRequest, session string) (*film.TaobaofilmtfbackyardcardscheduleupdateAPIResponse, error) {
	var resp film.TaobaofilmtfbackyardcardscheduleupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
