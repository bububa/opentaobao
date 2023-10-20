package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterServicestoreUpdateservicestorecoverservice 更新网点覆盖的服务
// tmall.servicecenter.servicestore.updateservicestorecoverservice
//
// 更新网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要新增的网点覆盖的服务不存在，会更新失败。
// 网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
func TmallServicecenterServicestoreUpdateservicestorecoverservice(clt *core.SDKClient, req *tmallsc.TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIRequest, resp *tmallsc.TmallServicecenterServicestoreUpdateservicestorecoverserviceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
