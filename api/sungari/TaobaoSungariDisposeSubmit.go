package sungari

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// Taobaosungaridisposesubmit 商品商家处置提交任务
// taobao.sungari.dispose.submit
//
// 商品商家处置信息接口，提供政府部门发送处置信息给阿里
func Taobaosungaridisposesubmit(clt *core.SDKClient, req *sungari.TaobaosungaridisposesubmitAPIRequest, session string) (*sungari.TaobaosungaridisposesubmitAPIResponse, error) {
	var resp sungari.TaobaosungaridisposesubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
