package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeThreedimensionSync 二手房3D户型信息同步
// alibaba.alihouse.existinghome.threedimension.sync
//
// 二手房3D户型信息同步
func AlibabaAlihouseExistinghomeThreedimensionSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeThreedimensionSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeThreedimensionSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeThreedimensionSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
