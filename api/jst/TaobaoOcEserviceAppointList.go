package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoOcEserviceAppointList 交互卡片预约信息读取接口
// taobao.oc.eservice.appoint.list
//
// 允许外部的isv通过该接口读取门店预约信息
func TaobaoOcEserviceAppointList(clt *core.SDKClient, req *jst.TaobaoOcEserviceAppointListAPIRequest, session string) (*jst.TaobaoOcEserviceAppointListAPIResponse, error) {
	var resp jst.TaobaoOcEserviceAppointListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
