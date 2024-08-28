package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// TaobaoWdkEquipmentWcsWcsinfoUpload 悬挂链业务信息上传
// taobao.wdk.equipment.wcs.wcsinfo.upload
//
// 五道口仓库悬挂链信息上传
func TaobaoWdkEquipmentWcsWcsinfoUpload(ctx context.Context, clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentWcsWcsinfoUploadAPIRequest, resp *wdk.TaobaoWdkEquipmentWcsWcsinfoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
