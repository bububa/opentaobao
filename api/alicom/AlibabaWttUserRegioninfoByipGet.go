package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabawttuserregioninfobyipget 根据ip获取省市信息
// alibaba.wtt.user.regioninfo.byip.get
//
// 通过ip获取省市信息
func Alibabawttuserregioninfobyipget(clt *core.SDKClient, req *alicom.AlibabawttuserregioninfobyipgetAPIRequest, session string) (*alicom.AlibabawttuserregioninfobyipgetAPIResponse, error) {
	var resp alicom.AlibabawttuserregioninfobyipgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
