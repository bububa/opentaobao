package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicestoreUpdateservicestore 修改网点信息
// tmall.servicecenter.servicestore.updateservicestore
//
// 修改网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点存在才能修改。
// 错误码
// 1, 服务商昵称无效
// 2, 缺少省份
// 3, 缺少城市
// 4, 缺少区域
// 5, 缺少详细地址
// 6, 传入地址不在标准地址库中，无法解析
// 7, 缺少网点编码
// 8, 缺少网点名称
// 9, 缺少网点电话
// 10, 网点已存在
// 11, 网点不存在
// 12, 系统错误
func TmallServicecenterServicestoreUpdateservicestore(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreUpdateservicestoreAPIRequest, resp *tmallservice.TmallServicecenterServicestoreUpdateservicestoreAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
