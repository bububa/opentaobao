package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenterservicestoreupdateservicestorecapacity 更新网点容量
// tmall.servicecenter.servicestore.updateservicestorecapacity
//
// 更新网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要更新的网点容量不存在，会更新失败。
// 网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
func Tmallservicecenterservicestoreupdateservicestorecapacity(clt *core.SDKClient, req *tmallsc.TmallservicecenterservicestoreupdateservicestorecapacityAPIRequest, session string) (*tmallsc.TmallservicecenterservicestoreupdateservicestorecapacityAPIResponse, error) {
	var resp tmallsc.TmallservicecenterservicestoreupdateservicestorecapacityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
