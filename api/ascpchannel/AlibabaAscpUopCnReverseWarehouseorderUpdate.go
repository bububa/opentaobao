package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuopcnreversewarehouseorderupdate 供应链中台逆向入库单修改服务
// alibaba.ascp.uop.cn.reverse.warehouseorder.update
//
// 供应链中台逆向入库单修改服务
func Alibabaascpuopcnreversewarehouseorderupdate(clt *core.SDKClient, req *ascpchannel.AlibabaascpuopcnreversewarehouseorderupdateAPIRequest, session string) (*ascpchannel.AlibabaascpuopcnreversewarehouseorderupdateAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuopcnreversewarehouseorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
