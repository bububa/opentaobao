package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaologisticsappletmodifydatasave 物流小程序修改物流信息回传接口
// taobao.logistics.applet.modifydata.save
//
// 物流小程序修改物流信息回传接口
func Taobaologisticsappletmodifydatasave(clt *core.SDKClient, req *mtopopen.TaobaologisticsappletmodifydatasaveAPIRequest, session string) (*mtopopen.TaobaologisticsappletmodifydatasaveAPIResponse, error) {
	var resp mtopopen.TaobaologisticsappletmodifydatasaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
