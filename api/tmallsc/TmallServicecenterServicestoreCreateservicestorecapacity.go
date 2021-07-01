package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

/* TmallServicecenterServicestoreCreateservicestorecapacity
新增网点容量
tmall.servicecenter.servicestore.createservicestorecapacity

新增网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
前提是网点要存在，
如果需要新增的网点容量已存在，会新增失败。
网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量 */
func TmallServicecenterServicestoreCreateservicestorecapacity(clt *core.SDKClient, req *tmallsc.TmallServicecenterServicestoreCreateservicestorecapacityAPIRequest, session string) (*tmallsc.TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse, error) {
	var resp tmallsc.TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
