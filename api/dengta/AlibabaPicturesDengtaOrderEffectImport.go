package dengta

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// AlibabaPicturesDengtaOrderEffectImport 天下秀订单数据导入
// alibaba.pictures.dengta.order.effect.import
//
// 提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统
func AlibabaPicturesDengtaOrderEffectImport(ctx context.Context, clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaOrderEffectImportAPIRequest, resp *dengta.AlibabaPicturesDengtaOrderEffectImportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
