package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// Taobaologisticsshunfengmodifydatasave 顺丰小程序修改配送信息回传接口
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
func Taobaologisticsshunfengmodifydatasave(clt *core.SDKClient, req *mtopopen.TaobaologisticsshunfengmodifydatasaveAPIRequest, session string) (*mtopopen.TaobaologisticsshunfengmodifydatasaveAPIResponse, error) {
	var resp mtopopen.TaobaologisticsshunfengmodifydatasaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
