package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// CainiaoCloudprintCustomareaUpdate 自定义区内容更新
// cainiao.cloudprint.customarea.update
//
// 自定义区内容更新
func CainiaoCloudprintCustomareaUpdate(clt *core.SDKClient, req *waybill.CainiaoCloudprintCustomareaUpdateAPIRequest, session string) (*waybill.CainiaoCloudprintCustomareaUpdateAPIResponse, error) {
	var resp waybill.CainiaoCloudprintCustomareaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
