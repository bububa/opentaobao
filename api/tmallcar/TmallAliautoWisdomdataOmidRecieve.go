package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautowisdomdataomidrecieve 大搜车车型参配数据接入
// tmall.aliauto.wisdomdata.omid.recieve
//
// 大搜车车型参配数据接入
func Tmallaliautowisdomdataomidrecieve(clt *core.SDKClient, req *tmallcar.TmallaliautowisdomdataomidrecieveAPIRequest, session string) (*tmallcar.TmallaliautowisdomdataomidrecieveAPIResponse, error) {
	var resp tmallcar.TmallaliautowisdomdataomidrecieveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
