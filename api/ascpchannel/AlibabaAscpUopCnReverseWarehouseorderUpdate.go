package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpUopCnReverseWarehouseorderUpdate 供应链中台逆向入库单修改服务
// alibaba.ascp.uop.cn.reverse.warehouseorder.update
//
// 供应链中台逆向入库单修改服务
func AlibabaAscpUopCnReverseWarehouseorderUpdate(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest, resp *ascpchannel.AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
