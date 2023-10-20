package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationGetInfoUuid 根据uuid用户抽奖次数限制
// taobao.degoperation.get.info.uuid
//
// 根据uuid用户抽奖次数限制
func TaobaoDegoperationGetInfoUuid(clt *core.SDKClient, req *degoperation.TaobaoDegoperationGetInfoUuidAPIRequest, session string) (*degoperation.TaobaoDegoperationGetInfoUuidAPIResponse, error) {
	var resp degoperation.TaobaoDegoperationGetInfoUuidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
