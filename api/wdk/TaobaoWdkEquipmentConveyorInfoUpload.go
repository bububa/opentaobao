package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Taobaowdkequipmentconveyorinfoupload 五道口仓库悬挂链信息上报
// taobao.wdk.equipment.conveyor.info.upload
//
// 五道口仓库悬挂链信息上传
func Taobaowdkequipmentconveyorinfoupload(clt *core.SDKClient, req *wdk.TaobaowdkequipmentconveyorinfouploadAPIRequest, session string) (*wdk.TaobaowdkequipmentconveyorinfouploadAPIResponse, error) {
	var resp wdk.TaobaowdkequipmentconveyorinfouploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
