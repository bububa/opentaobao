package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* TaobaoWdkEquipmentConveyorInfoUpload
五道口仓库悬挂链信息上报
taobao.wdk.equipment.conveyor.info.upload

五道口仓库悬挂链信息上传 */
func TaobaoWdkEquipmentConveyorInfoUpload(clt *core.SDKClient, req *wdk.TaobaoWdkEquipmentConveyorInfoUploadAPIRequest, session string) (*wdk.TaobaoWdkEquipmentConveyorInfoUploadAPIResponse, error) {
	var resp wdk.TaobaoWdkEquipmentConveyorInfoUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
